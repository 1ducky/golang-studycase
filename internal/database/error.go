package database

import (
	"context"
	"database/sql"
	"errors"
	"restApi/internal/application"
)

func mapDBError(err error) *application.AppError {
	switch {
	case errors.Is(err, sql.ErrConnDone):
		return &application.AppError{Code: "ERR_DB_CONN_CLOSED", Message: "Koneksi database sudah ditutup"}
	case errors.Is(err, context.DeadlineExceeded):
		return &application.AppError{Code: "ERR_DB_TIMEOUT", Message: "Query database melebihi batas waktu"}
	case errors.Is(err, context.Canceled):
		return &application.AppError{Code: "ERR_DB_CTX_CANCELED", Message: "Permintaan dibatalkan"}
	default:
		// bisa dicek juga pakai driver-specific error (mis. *pq.Error atau *mysql.MySQLError)
		return &application.AppError{Code: "ERR_DB_UNKNOWN", Message: "Terjadi kesalahan pada database"}
	}
}
