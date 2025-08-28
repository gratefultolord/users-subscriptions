package update_subscription

import (
	"context"
	"fmt"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
)

type Usecase struct {
	storage storage
}

func NewUsecase(storage storage) *Usecase {
	return &Usecase{
		storage: storage,
	}
}

func (u *Usecase) UpdateSubscriptionInfo(ctx context.Context, subscription domain.Subscription) error {
	if err := u.storage.UpdateSubscriptionInfo(ctx, subscription); err != nil {
		return fmt.Errorf("u.storage.UpdateSubscriptionInfo: %w", err)
	}

	return nil
}
