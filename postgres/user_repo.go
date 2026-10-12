package postgres

import (
	"context"
	"twitter"
)

type UserRepo struct {
	DB *DB
}

func (u UserRepo) GetByUsername(ctx context.Context, username string) (*twitter.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) GetByEmail(ctx context.Context, email string) (*twitter.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Create(ctx context.Context, user twitter.User) (twitter.User, error) {
	//TODO implement me
	panic("implement me")
}
