package database

import (
	"log"

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
		log.Fatal(err)
	}

	err = m.Up()

	if err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	log.Println("Migrations applied")
}