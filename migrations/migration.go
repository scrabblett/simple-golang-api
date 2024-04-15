package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
	"go.uber.org/zap"
	"strings"
)

const migrationsRoot = "migrations/scripts"

func RunMigrations(db *sql.DB) {
	err := goose.Up(db, migrationsRoot)

	if err != nil && !strings.Contains(err.Error(), "no migrations to run") {
		zap.L().Fatal("Failed to run migrations", zap.Error(err))
	}
}
