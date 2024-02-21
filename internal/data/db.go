package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

// get env vars and build db url
func NewDB(ctx context.Context, dbUrl string) (*DB, error) { // env.MustGet("DB_URL")
	pool, err := pgxpool.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}
	return &DB{pool: pool}, nil
}

// # Example DSN
// user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca

// # Example URL
// postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca
