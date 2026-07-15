package config

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/vsayfb/gig-platform-notification-lambda/pkg/fb"
)

func loadEnv() (*Config, error) {
	return &Config{
		APP: AppConfig{
			Env:                     getEnv("APP_ENV", "development"),
			FirebaseCredentialsPath: mustGetEnv("FIREBASE_CREDENTIALS_PATH"),
		},
		DB: DBConfig{},
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

func (c *AppConfig) GetFireBaseCredentials() (*fb.FirebaseServiceAccount, error) {

	if _, err := os.Stat(c.FirebaseCredentialsPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("fcm: credentials file not found at %s: %w", c.FirebaseCredentialsPath, err)
	}

	credBytes, err := os.ReadFile(c.FirebaseCredentialsPath)

	if err != nil {
		return nil, fmt.Errorf("fcm: failed to read credentials: %w", err)
	}

	slog.Info("Credentials file read successfully", "size_bytes", len(credBytes))

	var credentials fb.FirebaseServiceAccount

	if err := json.Unmarshal(credBytes, &credentials); err != nil {
		return nil, fmt.Errorf("fcm: failed to parse credentials JSON: %w", err)
	}

	return &credentials, nil
}
