package config

import (
	"os"
	"time"
)

type Config struct {
	HTTPAddr        string
	DatabaseURL     string
	JWTSecret       string
	ShutdownTimeout time.Duration
}

func Load() (Config, error) {
	return Config{
		HTTPAddr:        getEnv("HTTP_ADDR", ":8080"),
		DatabaseURL:     getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/subscription_service?sslmode=disable"),
		JWTSecret:       getEnv("JWT_SECRET", "dev-secret"),
		ShutdownTimeout: 10 * time.Second,
	}, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
