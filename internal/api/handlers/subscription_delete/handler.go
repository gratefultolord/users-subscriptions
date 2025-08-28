package subscription_delete

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
//	@Summary		Удалить подписку
//	@Description	Удаляет подписку по ID
//	@Tags			subscriptions
//	@Param			subscriptionId	path		int								true	"ID подписки"
//	@Success		204				{object}	response.DeletedResponse		"Success"
//	@Failure		400				{object}	response.BadRequestError		"Bad Request"
//	@Failure		404				{object}	response.NotFoundError			"Not Found"
//	@Failure		500				{object}	response.InternalServerError	"Internal Server Error"
//	@Router			/subscriptions/{subscriptionId}/delete [delete]
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

	if err = h.usecase.DeleteSubscriptionInfo(c.Request.Context(), subscriptionId); err != nil {
		if errors.Is(err, errs.ErrSubscriptionNotFound) {
			h.logger.Info("subscription not found",
				zap.Int64("subscriptionId", subscriptionId),
			)
			c.JSON(http.StatusNotFound, response.NotFoundError{Error: "Not Found"})
			return
		}

		h.logger.Error("h.usecase.DeleteSubscriptionInfo",
			zap.Int64("subscriptionId", subscriptionId),
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, response.InternalServerError{Error: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusNoContent, response.DeletedResponse{Result: "deleted"})
}
