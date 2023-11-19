package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

func GetConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
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
