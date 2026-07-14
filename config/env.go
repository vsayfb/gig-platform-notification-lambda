package config

import (
	"fmt"
	"os"
)

func loadEnv() (*Config, error) {
	return &Config{
		APP: AppConfig{
			Env:                     getEnv("APP_ENV", "development"),
			FirebaseCredentialsPath: mustGetEnv("FIREBASE_CREDENTIALS_PATH"),
		},
		DB: DBConfig{
			Host:     mustGetEnv("DB_HOST"),
			Port:     mustGetEnv("DB_PORT"),
			User:     mustGetEnv("DB_USER"),
			Password: mustGetEnv("DB_PASSWORD"),
			Name:     mustGetEnv("DB_NAME"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		TEL: TelemetryConfig{
			getEnv("OTEL_COLLECTOR_ADDR", "localhost:4317"),
		},
	}, nil
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)

	if v == "" {
		panic(fmt.Sprintf("missing required env var: %s", key))
	}

	return v
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
