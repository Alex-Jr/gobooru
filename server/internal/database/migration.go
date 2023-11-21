package database

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

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
	_, curPath, _, _ := runtime.Caller(0)
	migrationPath := filepath.Join(filepath.Dir(curPath), "../../migrations")

	driver, err := postgres.WithInstance(
		db.DB,
		&postgres.Config{},
	)
	if err != nil {
		log.Fatalf("error creating migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		// TODO: using env because docker, tests and debug have different paths
		fmt.Sprintf("file:///%s", migrationPath),
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

	log.Println("migrations ran successfully")

	return nil
}
