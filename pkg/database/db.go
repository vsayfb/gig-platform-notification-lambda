package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, databaseURL)

	if err != nil {
		return nil, fmt.Errorf("database: failed to create pool: %w", err)
	}

	slog.Info("db connected")

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("database: ping failed: %w", err)
	}

	return pool, nil
}
