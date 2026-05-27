package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/raystack/optimus/config"
)

// Open will connect to the DB with custom configuration
func Open(config config.DBConfig) (*pgxpool.Pool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cleanup to be done with dbPool.Close()
