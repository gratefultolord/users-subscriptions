package get_subscriptions_list

import (
	"context"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
)

type storage interface {
	GetSubscriptions(ctx context.Context, limit, offset int) ([]domain.Subscription, error)
}
