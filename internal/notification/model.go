package notification

import (
	"time"

	"github.com/google/uuid"
)

type NotificationType string

const (
	TypeGigMatch NotificationType = "gig_match"
)

type Notification struct {
	ID        uuid.UUID        `json:"id"`
	UserID    uuid.UUID        `json:"user_id"`
	Type      NotificationType `json:"type"`
	RefGigID  *uuid.UUID       `json:"ref_gig_id,omitempty"`
	Title     string           `json:"title"`
	Body      string           `json:"body"`
	IsRead    bool             `json:"is_read"`
	CreatedAt time.Time        `json:"created_at"`
}

type GigMatchMessage struct {
	GigID       uuid.UUID `json:"gig_id"`
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FCMToken    string    `json:"fcm_token"`
}
