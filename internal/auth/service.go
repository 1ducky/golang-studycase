package auth

import (
	"context"
	"restApi/internal/auth/jwt"
)

type Service struct {
	// method
	JWT *jwt.Service

	Repo Repository
}

func NewAuthService(repo Repository, jwt *jwt.Service) *Service {
	return &Service{JWT: jwt, Repo: repo}
}

func (s *Service) Login(ctx context.Context, payload LoginRequest) (string, error) {
	if payload.Password == "" || payload.Username == "" {
		return "", ErrorBadRequest
	}

	user, err := s.Repo.Login(ctx, payload)
	if err != nil {
		return "", err
	}
	jwtToken, err := s.JWT.CreateAuthClaim(user.ID, user.Username, user.Role)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (s *Service) ParseJWTToken(ctx context.Context, token string) (AuthClaims, error) {
	user, err := s.JWT.ParseAuthClaim(token)
	if err != nil {
		return AuthClaims{}, err
	}
	return AuthClaims{ID: user.ID, Username: user.Username, Role: Role(user.Role)}, nil
}
