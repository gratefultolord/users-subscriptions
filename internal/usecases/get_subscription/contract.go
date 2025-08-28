package get_subscription

import (
	"context"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
)

type storage interface {
	GetSubscription(ctx context.Context, subscriptionId int64) (domain.Subscription, error)
}
