package main

import (
	"context"
	"log"

	"github.com/vsayfb/gig-platform-notification-lambda/config"
	"github.com/vsayfb/gig-platform-notification-lambda/pkg/database"
	"github.com/vsayfb/gig-platform-notification-lambda/pkg/fb"
	"github.com/vsayfb/gig-platform-notification-lambda/pkg/logger"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load(ctx)

	if err != nil {
		log.Fatal(err)
	}

	logger.Init(cfg.APP.Env)

	_, err = database.NewPool(ctx, cfg.DB.DSN())

	if err != nil {
		log.Fatal(err)
	}

	creds, err := cfg.APP.GetFireBaseCredentials()

	if err != nil {
		log.Fatal(err)
	}

	_, err = fb.NewClient(ctx, creds)
}
