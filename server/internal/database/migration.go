package database

import (
	"log"

	"github.com/go-sqlx/sqlx"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MustRunMigrations(db *sqlx.DB) {
	err := RunMigrations(db)
	if err != nil {
		log.Fatalf("error running migrations: %v", err)
	}
}

func RunMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(
		db.DB,
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

	return nil
}
