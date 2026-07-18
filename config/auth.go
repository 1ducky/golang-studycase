package config

import (
	"os"
)

type AuthConfig struct {
	Key string
}

func NewAuthConfig() *AuthConfig {
	return &AuthConfig{
		Key: os.Getenv("SECRET_AUTH"),
	}
}
