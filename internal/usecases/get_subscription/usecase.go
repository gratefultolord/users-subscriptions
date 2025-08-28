package get_subscription

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

func (u *Usecase) GetSubscription(ctx context.Context, subscriptionID int64) (domain.Subscription, error) {
	subscription, err := u.storage.GetSubscription(ctx, subscriptionID)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("u.storage.GetSubscription:  %w", err)
	}
	return subscription, nil
}
