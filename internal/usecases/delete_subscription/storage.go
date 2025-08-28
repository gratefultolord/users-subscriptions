package delete_subscription

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (s *Storage) DeleteSubscriptionInfo(ctx context.Context, subscriptionID int64) error {
	query := `
		DELETE FROM subscriptions
		WHERE id = $1
		RETURNING id
	`

	var deletedID int64
	err := s.db.QueryRowxContext(ctx, query, subscriptionID).Scan(&deletedID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrSubscriptionNotFound
		}

		return fmt.Errorf("s.db.QueryRowxContext: %w", err)
	}

	return nil
}
