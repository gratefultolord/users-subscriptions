package delete_subscription

import (
	"context"
	"fmt"
)

type Usecase struct {
	storage storage
}

func NewUsecase(storage storage) *Usecase {
	return &Usecase{storage: storage}
}

func (u *Usecase) DeleteSubscriptionInfo(ctx context.Context, subscriptionID int64) error {
	if err := u.storage.DeleteSubscriptionInfo(ctx, subscriptionID); err != nil {
		return fmt.Errorf("u.storage.DeleteSubscriptionInfo: %w", err)
	}

	return nil
}
