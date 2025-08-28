package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          int64      `db:"id"`
	ServiceName string     `db:"service_name"`
	Price       int64      `db:"price"`
	UserID      uuid.UUID  `db:"user_id"`
	StartDate   time.Time  `db:"start_date"`
	EndDate     *time.Time `db:"end_date"`
}

type SubscriptionDTO struct {
	ID          int64   `json:"id"`
	ServiceName string  `json:"service_name"`
	Price       int64   `json:"price"`
	UserID      string  `json:"user_id"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date,omitempty"`
}
