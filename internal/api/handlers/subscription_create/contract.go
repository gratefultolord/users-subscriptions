package subscription_create

import (
	"context"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
)

type usecase interface {
	CreateSubscriptionInfo(ctx context.Context, subscription domain.Subscription) error
}
