package get_total_price

import (
	"context"

	"github.com/google/uuid"
)

type storage interface {
	GetTotalPrice(ctx context.Context, userID *uuid.UUID, serviceName *string) (*int64, error)
}
