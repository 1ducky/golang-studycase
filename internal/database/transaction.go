package database

import (
	"context"
	"database/sql"
)

// DBTX adalah interface yang dipenuhi baik oleh *sql.DB maupun *sql.Tx,
// sehingga kode/query bisa dijalankan dalam transaksi atau tanpa transaksi.
type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type TransactionManager struct {
	DB *sql.DB
}

func NewDBTransaction(Db *sql.DB) *TransactionManager {
	return &TransactionManager{DB: Db}
}

func (t *TransactionManager) Do(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	err = fn(tx)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
