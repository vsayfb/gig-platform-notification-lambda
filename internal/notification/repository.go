package notification

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const insertNotificationQuery = `
INSERT INTO notifications (user_id, type, ref_gig_id, title, body, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
`

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Insert(ctx context.Context, n *Notification) error {
	_, err := r.db.Exec(
		ctx,
		insertNotificationQuery,
		n.UserID,
		n.Type,
		n.RefGigID,
		n.Title,
		n.Body,
		time.Now(),
	)
	
	if err != nil {
		return fmt.Errorf("insert notification: %w", err)
	}

	return nil
}
