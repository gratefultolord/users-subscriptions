package create_subscription

import (
	"context"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
)

type storage interface {
	AddSubscriptionInfo(ctx context.Context, subscription domain.Subscription) error
}
