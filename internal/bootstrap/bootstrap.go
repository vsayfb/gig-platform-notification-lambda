package bootstrap

import (
	"context"
	"fmt"

	"github.com/vsayfb/gig-platform-notification-lambda/config"
	"github.com/vsayfb/gig-platform-notification-lambda/internal/handler"
	"github.com/vsayfb/gig-platform-notification-lambda/internal/notification"
	"github.com/vsayfb/gig-platform-notification-lambda/pkg/database"
	"github.com/vsayfb/gig-platform-notification-lambda/pkg/fb"
	"github.com/vsayfb/gig-platform-notification-lambda/pkg/logger"
)

func NewHandler() (*handler.Handler, error) {
	ctx := context.Background()

	cfg, err := config.Load(ctx)

	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	logger.Init(cfg.APP.Env)

	db, err := database.NewPool(ctx, cfg.DB.DSN())

	if err != nil {
		return nil, fmt.Errorf("create database pool: %w", err)
	}

	creds, err := cfg.APP.GetFireBaseCredentials()

	if err != nil {
		return nil, fmt.Errorf("load firebase credentials: %w", err)
	}

	fbClient, err := fb.NewClient(ctx, creds)

	if err != nil {
		return nil, fmt.Errorf("create firebase client: %w", err)
	}

	repo := notification.NewRepository(db)

	service := notification.NewService(repo, fbClient)

	return handler.New(service), nil
}
