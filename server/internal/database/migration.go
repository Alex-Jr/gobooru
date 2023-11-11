package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MustRunMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(
		db,
		&postgres.Config{},
	)
	if err != nil {
		log.Fatalf("error creating migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///home/alex/projetos/gobooru/server/migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatalf("error creating migration: %v", err)
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("error running migration: %v", err)
	}
}
