package user

import (
	"context"
	"database/sql"
	"restApi/internal/utils"
)

type Repository interface {
	create(ctx context.Context, payload CreateRequest) (bool, error)
}

type MemoryReposiotry struct {
	DB *sql.DB
}

func NewUserMemory(Db *sql.DB) Repository {
	return &MemoryReposiotry{DB: Db}
}

func (r *MemoryReposiotry) create(ctx context.Context, payload CreateRequest) (bool, error) {
	query := "INSERT INTO user( `username`, `password`, `role`) VALUES (?,?,?)"
	hash, err := utils.HashPassword(payload.Password)
	if err != nil {
		return false, err
	}
	result, err := r.DB.ExecContext(ctx, query, payload.Username, hash, RoleUser)
	if err != nil {
		return false, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return false, err
	}

	return true, nil
}
