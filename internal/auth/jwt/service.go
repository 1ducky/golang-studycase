package jwt

import (
	"fmt"
	"restApi/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type Service struct {
	key string
}

func NewJWT(config *config.AuthConfig) *Service {
	return &Service{key: config.Key}
}

func (r *Service) CreateAuthClaim(id int, username string, role string) (string, error) {
	var secretKey = []byte(r.key)

	claims := AuthClaims{
		ID:       id,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "TodoApp",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (r *Service) ParseAuthClaim(tokenString string) (*AuthClaims, error) {
	claims := &AuthClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// pastikan signing method sesuai (best practice: cegah algorithm confusion attack)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(r.key), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
