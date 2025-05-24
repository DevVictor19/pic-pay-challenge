package env

import (
	"errors"
	"os"
)

type Config struct {
	ServerPort string
	DB         PostgresConfig
}

type PostgresConfig struct {
	URL string
}

var cfg *Config

func GetEnv() (*Config, error) {
	if cfg == nil {
		return nil, errors.New("config not initialized, call LoadEnv first")
	}

	return cfg, nil
}

func LoadEnv() (*Config, error) {
	cfg = &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		DB: PostgresConfig{
			URL: os.Getenv("DB_URL"),
		},
	}

	return cfg, nil
}
