package subscription_read

import (
	"context"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
)

type usecase interface {
	GetSubscription(ctx context.Context, subscriptionID int64) (domain.Subscription, error)
}
