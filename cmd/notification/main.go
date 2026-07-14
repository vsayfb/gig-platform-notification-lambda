package main

import (
	"context"
	"log"

	"github.com/vsayfb/gig-platform-notification-lambda/config"
	"github.com/vsayfb/gig-platform-notification-lambda/pkg/database"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load(ctx)

	if err != nil {
		log.Fatal(err)
	}

	_, err = database.NewPool(ctx, cfg.DB.DSN())

	if err != nil {
		log.Fatal(err)
	}

}
