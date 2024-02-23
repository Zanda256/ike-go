package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func (db *DB) InsertSource() {

}

func (db *DB) InsertDownload() {

}

// # Example DSN
// user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca

// # Example URL
// postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca

const (
	uniqueViolation = "23505"
	undefinedTable  = "42P01"
)

// Set of error variables for CRUD operations.
var (
	ErrDBNotFound        = sql.ErrNoRows
	ErrDBDuplicatedEntry = errors.New("duplicated entry")
	ErrUndefinedTable    = errors.New("undefined table")
)

// Config is the required properties to use the database.
type Config struct {
	User       string
	Password   string
	Host       string
	Name       string
	DisableTLS bool
	Pool       PoolConfig
}

type PoolConfig struct {
	MaxConns              string
	MinConns              string
	MaxConnLifetime       string
	MaxConnIdleTime       string
	HealthCheckPeriod     string
	MaxConnLifetimeJitter string
}

// Open knows how to open a database connection based on the configuration.
func Open(ctx context.Context, cfg Config) (*DB, error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")
	q.Set("pool_max_conns", cfg.Pool.MaxConns)
	q.Set("pool_min_conns", cfg.Pool.MinConns)
	q.Set("pool_max_conn_lifetime", cfg.Pool.MaxConnLifetime)
	q.Set("pool_max_conn_idle_time", cfg.Pool.MaxConnIdleTime)
	q.Set("pool_health_check_period", cfg.Pool.HealthCheckPeriod)
	q.Set("pool_max_conn_lifetime_jitter", cfg.Pool.MaxConnLifetimeJitter)

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	dbPool, err := pgxpool.Connect(ctx, u.String())
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}
	return &DB{dbPool}, nil
}

// StatusCheck returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise.
func StatusCheck(ctx context.Context, db *DB) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second)
		defer cancel()
	}

	var pingError error
	for attempts := 1; ; attempts++ {
		pingError = db.Ping(ctx)
		if pingError == nil {
			break
		}
		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Run a simple query to determine connectivity.
	// Running this query forces a round trip through the database.
	const q = `SELECT true`
	var tmp bool
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	return tx.QueryRow(ctx, q, nil).Scan(&tmp)
}
