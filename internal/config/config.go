package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	HTTPAddr    string
	LogLevel    string

	DatabaseURL string
	RedisAddr   string
	NATSURL     string
}

func Load() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		Environment: getEnv("APP_ENV", "development"),
		HTTPAddr:    getEnv("APP_HTTP_ADDR", ":8080"),
		LogLevel:    getEnv("LOG_LEVEL", "debug"),

		DatabaseURL: getEnv("DATABASE_URL", ""),
		RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
		NATSURL:     getEnv("NATS_URL", "nats://localhost:4222"),
	}

	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
