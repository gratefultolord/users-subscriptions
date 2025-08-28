package subscription_read

import (
	"github.com/AlekSi/pointer"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/response"
)

func Present(subscription domain.Subscription) response.SubscriptionResponse {
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
}
