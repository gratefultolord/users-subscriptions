package subscription_list

import (
	"context"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
)

type usecase interface {
	GetSubscriptions(ctx context.Context, limit, offset int) ([]domain.Subscription, error)
}
