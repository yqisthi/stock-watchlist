package database

import (
	"stock-watchlist/infrastructure/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5432/stock_watchlist?sslmode=disable",
	)

	if err != nil {
		logger.Error(err, "Failed to initialize migrations")
	}

	err = m.Up()

	if err != nil && err.Error() != "no change" {
		logger.Error(err, "Failed to run migrations")
	}

	logger.Info("Migrations applied")
}