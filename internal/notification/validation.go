package notification

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

func validateNotification(n *Notification) error {
	switch {
	case n == nil:
		return errors.New("notification is nil")

	case n.UserID == uuid.Nil:
		return errors.New("user id is required")

	case strings.TrimSpace(n.Title) == "":
		return errors.New("title is required")

	case strings.TrimSpace(n.Body) == "":
		return errors.New("body is required")

	default:
		return nil
	}
}
