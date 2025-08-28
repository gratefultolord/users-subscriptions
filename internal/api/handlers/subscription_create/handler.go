package subscription_create

import (
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/request"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/response"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/utils"
	"go.uber.org/zap"
)

type Handler struct {
	logger  *zap.Logger
	usecase usecase
}

func NewHandler(logger *zap.Logger, usecase usecase) *Handler {
	return &Handler{
		logger:  logger,
		usecase: usecase,
	}
}

// Handle godoc
//
//	@Summary		Добавить информацию о подписке
//	@Description	Добавляет информацию о подписке
//	@Tags			subscriptions
//	@Accept			json
//	@Produce		json
//	@Param			input	body		request.SubscriptionRequest		true	"Данные подписки"
//	@Success		200		{object}	response.CreatedResponse		"Success"
//	@Failure		400		{object}	response.BadRequestError		"Bad Request"
//	@Failure		500		{object}	response.InternalServerError	"Internal Server Error"
//	@Router			/subscriptions/create [post]
func (h *Handler) Handle(c *gin.Context) {
	var req request.SubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("invalid request input",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, response.BadRequestError{Error: "Bad Request"})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		h.logger.Warn("could not parse userID",
			zap.String("userID", req.UserID),
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, response.BadRequestError{Error: "Bad Request"})
		return
	}

	if err := h.usecase.CreateSubscriptionInfo(c.Request.Context(), domain.Subscription{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      userID,
		StartDate:   pointer.Get(utils.MonthYearToTime(pointer.To(req.StartDate))),
		EndDate:     utils.MonthYearToTime(req.EndDate),
	}); err != nil {
		h.logger.Error("h.usecase.CreateSubscriptionInfo: %v",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, response.InternalServerError{Error: "Internal Server Error"})

		return
	}

	c.JSON(http.StatusCreated, response.CreatedResponse{Result: "created"})
}
