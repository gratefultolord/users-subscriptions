package total_price_get

import (
	"context"

	"github.com/google/uuid"
)

type usecase interface {
	GetTotalPrice(ctx context.Context, userID *uuid.UUID, serviceName *string) (*int64, error)
}
