package request

type SubscriptionRequest struct {
	ServiceName string  `json:"service_name" binding:"required" example:"Yandex Plus"`
	Price       int64   `json:"price" binding:"required" example:"1000"`
	UserID      string  `json:"user_id" binding:"required" example:"bac8ff49-1681-445c-941b-c000f2fc8ac0"`
	StartDate   string  `json:"start_date" binding:"required" example:"06-2025"`
	EndDate     *string `json:"end_date,omitempty" example:"12-2025" nullable:"true"`
}
