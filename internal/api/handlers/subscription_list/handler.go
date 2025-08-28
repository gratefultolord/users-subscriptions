package subscription_list

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
//	@Summary		Получить список подписок
//	@Description	Возвращает список подписок с пагинацией
//	@Tags			subscriptions
//	@Param			limit	query		int									false	"Количество элементов на странице (по умолчанию 10)"
//	@Param			offset	query		int									false	"Смещение (по умолчанию 0)"
//	@Success		200		{object}	response.SubscriptionListResponse	"Success"
//	@Failure		400		{object}	response.BadRequestError			"Bad Request"
//	@Failure		500		{object}	response.InternalServerError		"Internal Server Error"
//	@Router			/subscriptions [get]
func (h *Handler) Handle(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		h.logger.Warn("invalid limit",
			zap.String("limit", limitStr),
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, response.BadRequestError{Error: "Bad Request"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		h.logger.Warn("invalid offset",
			zap.String("offset", offsetStr),
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, response.BadRequestError{Error: "Bad Request"})
		return
	}

	subscriptions, err := h.usecase.GetSubscriptions(c.Request.Context(), limit, offset)
	if err != nil {
		h.logger.Error("h.usecase.GetSubscriptions",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, response.InternalServerError{Error: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, response.SubscriptionListResponse{Subscriptions: Present(subscriptions)})
}
