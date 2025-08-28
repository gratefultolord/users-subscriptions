package subscription_read

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/errs"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/response"
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
//	@Summary		Получить подписку по ID
//	@Description	Возвращает данные одной подписки по её идентификатору
//	@Tags			subscriptions
//	@Param			subscriptionId	path		int								true	"ID подписки"
//	@Success		200				{object}	response.SubscriptionView		"Success"
//	@Failure		400				{object}	response.BadRequestError		"Bad Request"
//	@Failure		404				{object}	response.NotFoundError			"Not Found"
//	@Failure		500				{object}	response.InternalServerError	"Internal Server Error"
//	@Router			/subscriptions/{subscriptionId} [get]
func (h *Handler) Handle(c *gin.Context) {
	subscriptionIdStr := c.Param("subscriptionId")
	subscriptionId, err := strconv.ParseInt(subscriptionIdStr, 10, 64)
	if err != nil {
		h.logger.Warn("invalid subscriptionId",
			zap.String("subscriptionId", subscriptionIdStr),
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, response.BadRequestError{Error: "Bad Request"})
		return
	}

	subscription, err := h.usecase.GetSubscription(c.Request.Context(), subscriptionId)
	if err != nil {
		if errors.Is(err, errs.ErrSubscriptionNotFound) {
			h.logger.Info("subscription not found",
				zap.Int64("subscriptionId", subscriptionId),
			)
			c.JSON(http.StatusNotFound, response.NotFoundError{Error: "Not Found"})
			return
		}

		h.logger.Error("h.usecase.GetSubscription",
			zap.Int64("subscriptionId", subscriptionId),
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, response.InternalServerError{Error: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, response.SubscriptionView{Subscription: Present(subscription)})
}
