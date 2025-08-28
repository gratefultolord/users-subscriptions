package get_subscriptions_list

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

func (s *Storage) GetSubscriptions(ctx context.Context, limit, offset int) ([]domain.Subscription, error) {
	subscriptions := make([]domain.Subscription, 0, limit)

	query := `SELECT id, service_name, price, user_id, start_date, end_date
	 		  FROM subscriptions 
			  LIMIT $1 OFFSET $2`

	err := s.db.SelectContext(ctx, &subscriptions, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("s.db.SelectContext: %w", err)
	}

	return subscriptions, nil
}
