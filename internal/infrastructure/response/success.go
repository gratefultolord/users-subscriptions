package response

type TotalPriceResponse struct {
	TotalPrice int64 `json:"total_price" example:"12345"`
}

type CreatedResponse struct {
	Result string `json:"result" example:"created"`
}

type UpdatedResponse struct {
	Result string `json:"result" example:"updated"`
}

type DeletedResponse struct {
	Result string `json:"result" example:"deleted"`
}

type SubscriptionResponse struct {
	ID          int64   `json:"id" example:"1234"`
	ServiceName string  `json:"service_name" example:"Netflix"`
	Price       int64   `json:"price" example:"1200"`
	UserID      string  `json:"user_id" example:"bac8ff49-1681-445c-941b-c000f2fc8ac0"`
	StartDate   string  `json:"start_date" example:"02-2025"`
	EndDate     *string `json:"end_date,omitempty" example:"03-2025"`
}

type SubscriptionListResponse struct {
	Subscriptions []SubscriptionResponse `json:"subcriptions"`
}

type SubscriptionView struct {
	Subscription SubscriptionResponse `json:"subscription"`
}
