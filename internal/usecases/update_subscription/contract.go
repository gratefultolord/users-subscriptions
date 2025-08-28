package update_subscription

import (
	"context"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
)

type storage interface {
	UpdateSubscriptionInfo(ctx context.Context, subscription domain.Subscription) error
}
