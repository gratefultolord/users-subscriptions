package get_total_price

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
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

func (s *Storage) GetTotalPrice(ctx context.Context, userID *uuid.UUID, serviceName *string) (*int64, error) {
	var total int64
	query := `
		SELECT SUM(price)
		FROM subscriptions
		WHERE ($1::uuid IS NULL OR user_id = $1)
  		  AND ($2::text IS NULL OR service_name = $2)
	`

	err := s.db.GetContext(ctx, &total, query, userID, serviceName)
	if err != nil {
		return nil, fmt.Errorf("s.db.GetContext: %w", err)
	}

	return pointer.To(total), nil
}
