package get_subscriptions_list

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

func (u *Usecase) GetSubscriptions(ctx context.Context, limit, offset int) ([]domain.Subscription, error) {
	subscriptions, err := u.storage.GetSubscriptions(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("u.storage.GetSubscriptions: %w", err)
	}

	return subscriptions, nil
}
