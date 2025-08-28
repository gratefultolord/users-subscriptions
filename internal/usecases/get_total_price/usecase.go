package get_total_price

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Usecase struct {
	storage storage
}

func NewUsecase(storage storage) *Usecase {
	return &Usecase{
		storage: storage,
	}
}

func (u *Usecase) GetTotalPrice(ctx context.Context, userID *uuid.UUID, serviceName *string) (*int64, error) {
	totalPrice, err := u.storage.GetTotalPrice(ctx, userID, serviceName)
	if err != nil {
		return nil, fmt.Errorf("u.storage.GetTotalPrice: %w", err)
	}

	return totalPrice, nil
}
