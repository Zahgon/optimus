package setup

import (
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbPool     *pgxpool.Pool
	initDBOnce sync.Once
)

func TestPool() *pgxpool.Pool { _ = "STUB: not implemented"; return nil }

func mustReadDBConfig() string { _ = "STUB: not implemented"; return "" }

// Did not find a suitable way to read db config

// migrateDB takes around 700ms to drop and recreate db + run migrations
func migrateDB() { _ = "STUB: not implemented"; return }

func cleanDB(m *migrate.Migrate, pool *pgxpool.Pool) { _ = "STUB: not implemented"; return }

func dropTables(db *pgxpool.Pool) error { _ = "STUB: not implemented"; return nil }

func TruncateTablesWith(pool *pgxpool.Pool) { _ = "STUB: not implemented"; return }
