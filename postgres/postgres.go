package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path"
	"runtime"
	"twitter/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Pool   *pgxpool.Pool
	config *config.Config
}

func New(ctx context.Context, conf *config.Config) *DB {
	log.Println("url===>", conf.Database.URL)
	dbConf, err := pgxpool.ParseConfig(conf.Database.URL)
	if err != nil {
		log.Fatalf("failed to parse postgres connection string: %v", err)
	}
	conn, err := pgxpool.ConnectConfig(ctx, dbConf)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	db := &DB{
		Pool:   conn,
		config: conf,
	}
	db.Ping(ctx)
	return db
}
func (db *DB) Ping(ctx context.Context) {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	log.Println("Successfully connected to database")
}
func (db *DB) Close() {
	db.Pool.Close()
}

func (db *DB) Migrate() error {
	_, b, _, _ := runtime.Caller(0)
	migrationDir := fmt.Sprintf("file:///%s/migrations", path.Dir(b))
	m, err := migrate.New(migrationDir, db.config.Database.URL)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	log.Println("Successfully migrated database")
	return nil
}
