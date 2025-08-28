package delete_subscription

import "context"

type storage interface {
	DeleteSubscriptionInfo(ctx context.Context, subscriptionId int64) error
}
