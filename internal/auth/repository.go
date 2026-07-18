package auth

import (
	"context"
	"database/sql"
	"restApi/internal/utils"
)

type Repository interface {
	Login(ctx context.Context, payload LoginRequest) (LoginData, error)
}

type AuthMemory struct {
	DB *sql.DB
}

func NewRepository(Db *sql.DB) Repository {
	return &AuthMemory{DB: Db}
}

func (r *AuthMemory) Login(ctx context.Context, payload LoginRequest) (LoginData, error) {
	var temp tempUserAuth
	query := "SELECT id,username,password,role FROM user WHERE username = ?"
	err := r.DB.QueryRowContext(ctx, query, payload.Username).Scan(
		&temp.ID,
		&temp.Username,
		&temp.password,
		&temp.Role,
	)
	if err != nil {
		return LoginData{}, ErrorNotFound
	}
	ok := utils.CheckPasswordHash(payload.Password, temp.password)
	if !ok {
		return LoginData{}, ErrorBadRequest
	}

	return LoginData{ID: temp.ID, Username: temp.Username, Role: temp.Role}, nil
}
