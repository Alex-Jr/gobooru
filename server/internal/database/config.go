package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sqlx/sqlx"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func GetConnectionString(c DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
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
		Port:     "5450",
		User:     "user",
		Password: "password",
		Database: "database",
	}
}

func MustGetSQLXConnection(c DBConfig) *sqlx.DB {
	return sqlx.MustConnect(
		"postgres",
		GetConnectionString(c),
	)
}

func GetSQLXConnection(c DBConfig) (*sqlx.DB, error) {
	return sqlx.Connect(
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
