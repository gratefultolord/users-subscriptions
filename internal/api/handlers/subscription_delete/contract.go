package subscription_delete

import "context"

type usecase interface {
	DeleteSubscriptionInfo(ctx context.Context, subscriptionID int64) error
}
