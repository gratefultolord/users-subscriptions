package subscription_update

import (
	"context"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
)

type usecase interface {
	UpdateSubscriptionInfo(ctx context.Context, subscription domain.Subscription) error
}
