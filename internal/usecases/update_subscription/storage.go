package update_subscription

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/errs"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) UpdateSubscriptionInfo(ctx context.Context, subscription domain.Subscription) error {
	query := `
		UPDATE subscriptions
		SET 
			service_name = $1,
			price        = $2,
			user_id      = $3,
			start_date   = $4,
			end_date     = $5
		WHERE id = $6
		RETURNING id
	`

	var updatedID int64
	err := s.db.QueryRowContext(
		ctx,
		query,
		subscription.ServiceName,
		subscription.Price,
		subscription.UserID,
		subscription.StartDate,
		subscription.EndDate,
		subscription.ID,
	).Scan(&updatedID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.ErrSubscriptionNotFound
		}
		return fmt.Errorf("s.db.QueryRowContext: %w", err)
	}

	return nil
}
