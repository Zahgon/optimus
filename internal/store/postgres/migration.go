package postgres

import (
	"embed"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // required for postgres migrate driver
)

//go:embed migrations
var migrationFs embed.FS

const (
	resourcePath = "migrations"
)

func NewMigrator(dbConnURL string) (*migrate.Migrate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Migrate to run up migrations
func Migrate(connURL string) error { _ = "STUB: not implemented"; return nil }

// Rollback to run up migrations
func Rollback(connURL string, count int) error { _ = "STUB: not implemented"; return nil }

func ToVersion(version uint, connURL string) error { _ = "STUB: not implemented"; return nil }
