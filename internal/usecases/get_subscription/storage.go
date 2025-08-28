package get_subscription

import (
	"context"
	"database/sql"
	"errors"
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

func (s *Storage) GetSubscription(ctx context.Context, subscriptionID int64) (domain.Subscription, error) {
	var subscription domain.Subscription
	query := `
		SELECT id, service_name, price, user_id, start_date, end_date
		FROM subscriptions
		WHERE id = $1
	`

	err := s.db.GetContext(ctx, &subscription, query, subscriptionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Subscription{}, errs.ErrSubscriptionNotFound
		}

		return domain.Subscription{}, fmt.Errorf("s.db.GetContext: %w", err)
	}

	return subscription, nil
}
