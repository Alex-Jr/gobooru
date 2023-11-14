package database

import (
	"context"
	"database/sql"

	"github.com/go-sqlx/sqlx"
)

type TransactionClient interface {
	DBClient
	Commit() error
	Rollback() error
}

type transactionClient struct {
	tx *sqlx.Tx
}

type TxOptions struct {
	*sql.TxOptions
}

func NewTransactionClient(tx *sqlx.Tx) TransactionClient {
	return &transactionClient{tx: tx}
}

func (c *transactionClient) Commit() error {
	return c.tx.Commit()
}

func (c *transactionClient) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.tx.ExecContext(ctx, query, args...)
}

func (c *transactionClient) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return c.tx.GetContext(ctx, dest, query, args...)
}

func (c *transactionClient) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	//? tx is missing export for NamedQueryContext
	return sqlx.NamedQueryContext(ctx, c.tx, query, arg)
}

func (c *transactionClient) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return c.tx.NamedExecContext(ctx, query, arg)
}

func (c *transactionClient) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return c.tx.SelectContext(ctx, dest, query, args...)
}

func (c *transactionClient) Rollback() error {
	return c.tx.Rollback()
}
