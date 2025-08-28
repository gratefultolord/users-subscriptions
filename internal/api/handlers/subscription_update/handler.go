package subscription_update

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/errs"
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
//	@Summary		Обновить подписку
//	@Description	Обновляет информацию о подписке по её ID
//	@Tags			subscriptions
//	@Accept			json
//	@Produce		json
//	@Param			subscriptionId	path		int								true	"ID подписки"
//	@Param			request			body		request.SubscriptionRequest		true	"Данные для обновления подписки"
//	@Success		200				{object}	response.UpdatedResponse		"Success"
//	@Failure		400				{object}	response.BadRequestError		"Bad Request"
//	@Failure		404				{object}	response.NotFoundError			"Not Found"
//	@Failure		500				{object}	response.InternalServerError	"Internal Server Error"
//	@Router			/subscriptions/{subscriptionId}/update [put]
func (h *Handler) Handle(c *gin.Context) {
	subscriptionIdStr := c.Param("subscriptionId")
	subscriptionId, err := strconv.ParseInt(subscriptionIdStr, 10, 64)
	if err != nil {
		h.logger.Warn("invalid subscriptionId",
			zap.String("subscriptionId", subscriptionIdStr),
		)
		c.JSON(http.StatusBadRequest, response.BadRequestError{Error: "Bad Request"})
		return
	}

	var req request.SubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("invalid request input",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, response.BadRequestError{Error: "Bad Request"})
		return
	}

	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		h.logger.Warn("could not parse userId",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, response.BadRequestError{Error: "Bad Request"})
		return
	}

	if err := h.usecase.UpdateSubscriptionInfo(c.Request.Context(), domain.Subscription{
		ID:          subscriptionId,
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      userUUID,
		StartDate:   pointer.Get(utils.MonthYearToTime(pointer.To(req.StartDate))),
		EndDate:     utils.MonthYearToTime(req.EndDate),
	}); err != nil {
		if errors.Is(err, errs.ErrSubscriptionNotFound) {
			h.logger.Info("subscription not found",
				zap.Int64("subscriptionId", subscriptionId),
			)
			c.JSON(http.StatusNotFound, response.NotFoundError{Error: "Not Found"})
			return
		}

		h.logger.Error("h.usecase.UpdateSubscriptionInfo",
			zap.Int64("subscriptionId", subscriptionId),
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, response.InternalServerError{Error: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, response.UpdatedResponse{Result: "updated"})
}
