package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/vsayfb/gig-platform-notification-lambda/internal/notification"
)

type Handler struct {
	service *notification.Service
}

func New(service *notification.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(ctx context.Context, event events.SQSEvent) error {
	var failed int

	for _, record := range event.Records {
		var msg notification.GigMatchMessage

		if err := json.Unmarshal([]byte(record.Body), &msg); err != nil {
			slog.Error(
				"Failed to unmarshal SQS message",
				"message_id", record.MessageId,
				"err", err,
			)

			failed++

			continue
		}

		if err := h.service.HandleGigMatch(ctx, msg); err != nil {
			slog.Error(
				"Failed to process notification",
				"user_id", msg.UserID,
				"gig_id", msg.GigID,
				"err", err,
			)

			failed++
		}
	}

	if failed > 0 {
		return fmt.Errorf("failed to process %d message(s)", failed)
	}

	return nil
}
