package notification

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/vsayfb/gig-platform-notification-lambda/pkg/fb"
)

type Service struct {
	repo      *Repository
	fcmClient *fb.FCMClient
}

func NewService(repo *Repository, fcmClient *fb.FCMClient) *Service {
	return &Service{
		repo:      repo,
		fcmClient: fcmClient,
	}
}

func (s *Service) HandleGigMatch(ctx context.Context, msg GigMatchMessage) error {
	slog.Info("Handling gig match",
		"gig_id", msg.GigID,
		"user_id", msg.UserID,
	)

	title, body := buildContent(msg)

	if err := s.persistNotification(ctx, msg, title, body); err != nil {
		return err
	}

	s.sendPush(ctx, msg, title, body)

	return nil
}

func buildContent(msg GigMatchMessage) (string, string) {
	body := msg.Description

	if len(body) > 255 {
		body = body[:252] + "..."
	}

	return msg.Title, body
}

func (s *Service) persistNotification(
	ctx context.Context,
	msg GigMatchMessage,
	title, body string,
) error {

	n := &Notification{
		UserID:   msg.UserID,
		Type:     TypeGigMatch,
		RefGigID: &msg.GigID,
		Title:    title,
		Body:     body,
		IsRead:   false,
	}

	if err := validateNotification(n); err != nil {
		return err
	}

	if err := s.repo.Insert(ctx, n); err != nil {
		slog.Error("Failed to insert notification", "err", err)

		return fmt.Errorf("notification service: failed to persist: %w", err)
	}

	slog.Info("Notification persisted successfully")

	return nil
}

func (s *Service) sendPush(
	ctx context.Context,
	msg GigMatchMessage,
	title, body string,
) {
	if s.fcmClient == nil {
		slog.Warn("FCM client not available, skipping push",
			"user_id", msg.UserID,
		)
		return
	}

	if msg.FCMToken == "" {
		slog.Warn("No FCM token, skipping push",
			"user_id", msg.UserID,
		)
		return
	}

	slog.Info("Sending FCM notification",
		"user_id", msg.UserID,
	)

	data := map[string]string{
		"type":   string(TypeGigMatch),
		"gig_id": msg.GigID.String(),
	}

	if err := s.fcmClient.Send(ctx, msg.FCMToken, title, body, data); err != nil {

		slog.Error("FCM push failed",
			"user_id", msg.UserID,
			"err", err,
		)

		return
	}

	slog.Info("FCM notification sent",
		"user_id", msg.UserID,
	)
}
