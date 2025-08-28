package create_subscription

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

func (u *Usecase) CreateSubscriptionInfo(ctx context.Context, subscription domain.Subscription) error {
	if err := u.storage.AddSubscriptionInfo(ctx, subscription); err != nil {
		return fmt.Errorf("u.storage.AddSubscriptionInfo: %w", err)
	}
	return nil
}
