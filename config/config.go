package config

import (
	"context"
	"fmt"
)

type Config struct {
	APP AppConfig
	DB  DBConfig
	TEL TelemetryConfig
}

type AppConfig struct {
	Env                     string
	FirebaseCredentialsPath string
}

type TelemetryConfig struct {
	OtelCollectorAddr string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load(ctx context.Context) (*Config, error) {
	return loadEnv()
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Name,
		c.SSLMode,
	)
}
