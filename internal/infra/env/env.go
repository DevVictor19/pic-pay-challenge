package env

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DB         PostgresConfig
}

type PostgresConfig struct {
	URL          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

var cfg *Config

func GetEnv() (*Config, error) {
	if cfg == nil {
		return nil, errors.New("config not initialized, call LoadEnv first")
	}

	return cfg, nil
}

func LoadEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg = &Config{
		ServerPort: getString("SERVER_PORT", "8080"),
		DB: PostgresConfig{
			URL:          getString("DB_URL", "postgres://admin:admin@localhost/picpay?sslmode=disable"),
			MaxOpenConns: getInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: getInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  getString("DB_MAX_IDLE_TIME", "15min"),
		},
	}

	return cfg, nil
}

func getString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func getInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}

// func getBool(key string, fallback bool) bool {
// 	val, ok := os.LookupEnv(key)
// 	if !ok {
// 		return fallback
// 	}

// 	boolVal, err := strconv.ParseBool(val)
// 	if err != nil {
// 		return fallback
// 	}

// 	return boolVal
// }
