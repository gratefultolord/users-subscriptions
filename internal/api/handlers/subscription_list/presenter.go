package subscription_list

import (
	"github.com/AlekSi/pointer"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/response"
	"github.com/samber/lo"
)

func Present(subscriptions []domain.Subscription) []response.SubscriptionResponse {

	return lo.Map(subscriptions, func(subscription domain.Subscription, _ int) response.SubscriptionResponse {
		dto := response.SubscriptionResponse{
			ID:          subscription.ID,
			ServiceName: subscription.ServiceName,
			Price:       subscription.Price,
			UserID:      subscription.UserID.String(),
			StartDate:   subscription.StartDate.Format("01-2006"),
		}

		if subscription.EndDate != nil {
			dto.EndDate = pointer.To(subscription.EndDate.Format("01-2006"))
		}

		return dto
	})
}
