package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sqlx/sqlx"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func GetConnectionString(c DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}

func GetDevConfig() DBConfig {
	return DBConfig{
		Host:     "localhost",
		Port:     5450,
		User:     "user",
		Password: "password",
		Database: "database",
	}
}

func GetSQLXConnection(c DBConfig) *sqlx.DB {
	return sqlx.MustConnect(
		"postgres",
		GetConnectionString(c),
	)
}

func GetSQLConnection(c DBConfig) *sql.DB {
	db, err := sql.Open(
		"postgres",
		GetConnectionString(c),
	)

	if err != nil {
		log.Fatalf("error opening connection: %v", err)
	}

	return db
}
