package config

type Config struct {
	Database *DatabaseConfig
	Auth     *AuthConfig
	Image    *ImageConfig
}

func NewConfig() *Config {
	return &Config{
		Database: NewDatabaseConfig(),
		Auth:     NewAuthConfig(),
		Image:    NewImageConfig(),
	}
}
