package config

type Config struct {
	Database *DatabaseConfig
	Auth     *AuthConfig
}

func NewConfig() *Config {
	return &Config{
		Database: NewDatabaseConfig(),
		Auth:     NewAuthConfig(),
	}
}
