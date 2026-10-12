package postgres

import (
	"context"
	"fmt"
	"twitter"
	errors2 "twitter/pkg/errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v4"
)

type UserRepo struct {
	DB *DB
}

func (ur UserRepo) GetByUsername(ctx context.Context, username string) (*twitter.User, error) {
	query := `SELECT * FROM users WHERE username = $1 LIMIT 1;`
	u := twitter.User{}
	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query, username); err != nil {
		if pgxscan.NotFound(err) {
			return &twitter.User{}, errors2.ErrNotFound
		}

		return &twitter.User{}, fmt.Errorf("error select: %w", err)
	}
	return &u, nil
}

func (ur UserRepo) GetByEmail(ctx context.Context, email string) (*twitter.User, error) {
	query := `select * FROM users WHERE email = $1 LIMIT 1;`
	u := twitter.User{}
	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query, email); err != nil {
		if pgxscan.NotFound(err) {
			return &twitter.User{}, errors2.ErrNotFound
		}

		return &twitter.User{}, fmt.Errorf("error select: %w", err)
	}
	return &u, nil
}

func (ur UserRepo) Create(ctx context.Context, user twitter.User) (twitter.User, error) {
	tx, err := ur.DB.Pool.Begin(ctx)
	if err != nil {
		return twitter.User{}, fmt.Errorf("error begin transaction: %w", err)
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)
	user, err = createUser(ctx, tx, user)
	if err != nil {
		return twitter.User{}, fmt.Errorf("error create user: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return twitter.User{}, fmt.Errorf("error commit transaction: %w", err)
	}
	return user, nil
}
func createUser(ctx context.Context, tx pgx.Tx, user twitter.User) (twitter.User, error) {
	query := `INSERT INTO users (username,email, password) VALUES ($1, $2, $3) RETURNING *;`
	u := twitter.User{}
	if err := pgxscan.Get(ctx, tx, &u, query, user.Username, user.Email, user); err != nil {
		return twitter.User{}, fmt.Errorf("error insert user: %w", err)
	}
	return u, nil
}
