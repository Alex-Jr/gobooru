package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-sqlx/sqlx"
)

type SQLClient interface {
	DBClient
	BeginTxx(ctx context.Context, opts *TxOptions) (TransactionClient, error)
}

type sqlClient struct {
	db *sqlx.DB
}

func NewSQLClient(db *sqlx.DB) SQLClient {
	return &sqlClient{
		db: db,
	}
}

func (c *sqlClient) BeginTxx(ctx context.Context, opts *TxOptions) (TransactionClient, error) {
	sqlOpts := sql.TxOptions{}

	if opts != nil {
		sqlOpts = *opts.TxOptions
	}

	tx, err := c.db.BeginTxx(ctx, &sqlOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	return NewTransactionClient(tx), nil
}

func (c *sqlClient) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.db.ExecContext(ctx, query, args...)
}

func (c *sqlClient) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return c.db.GetContext(ctx, dest, query, args...)
}

func (c *sqlClient) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	return c.db.NamedQueryContext(ctx, query, arg)
}

func (c *sqlClient) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return c.db.NamedExecContext(ctx, query, arg)
}

func (c *sqlClient) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return c.db.SelectContext(ctx, dest, query, args...)
}
