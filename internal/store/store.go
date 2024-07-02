package store

import (
	"context"
	"go-starter/internal/config"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type Store struct {
  DB *bun.DB
}

func New(db *bun.DB) *Store {
  return &Store{
    DB: db,
  }
}

func Open() *bun.DB {
	pool, err := pgxpool.New(context.Background(), config.MustEnv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}

	sqldb := stdlib.OpenDBFromPool(pool)
	if err := sqldb.Ping(); err != nil {
		log.Fatalf("unable to ping database: %v\n", err)
	}

	bun := bun.NewDB(sqldb, pgdialect.New())
  return bun
}
