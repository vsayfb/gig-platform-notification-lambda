package logger

import (
	"log/slog"
	"os"
)

const EnvProduction = "production"

func Init(env string) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	var handler slog.Handler

	switch env {
	case EnvProduction:
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(handler))
}
