package create_subscription

import (
	"context"
	"fmt"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
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

func (s *Storage) AddSubscriptionInfo(ctx context.Context, subscription domain.Subscription) error {
	query := `
		INSERT INTO subscriptions (
			service_name,
			price,
			user_id,
			start_date,
			end_date
		) VALUES (
		 	:service_name,
			:price,
			:user_id,
			:start_date,
			:end_date
		)
	`

	_, err := s.db.NamedExecContext(ctx, query, subscription)
	if err != nil {
		return fmt.Errorf("s.db.NamedExecContext: %w", err)
	}

	return nil
}
