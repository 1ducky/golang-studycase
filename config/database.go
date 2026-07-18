package config

import (
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}
}
