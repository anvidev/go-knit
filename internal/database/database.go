package database

import (
	"context"
	"go-starter/internal/config"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

// TODO: rename to store
// Testing
type Database struct {
	Postgres *bun.DB
}

func Open() *Database {
	pool, err := pgxpool.New(context.Background(), config.MustEnv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}

	sqldb := stdlib.OpenDBFromPool(pool)
	if err := sqldb.Ping(); err != nil {
		log.Fatalf("unable to ping database: %v\n", err)
	}

	bun := bun.NewDB(sqldb, pgdialect.New())
	return &Database{
		Postgres: bun,
	}
}
