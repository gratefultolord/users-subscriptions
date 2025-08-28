package total_price_get

import (
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/response"
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
//	@Summary		Получить общую стоимость подписок
//	@Description	Возвращает сумму всех подписок, опционально фильтруя по userId и serviceName
//	@Tags			subscriptions
//	@Accept			json
//	@Produce		json
//	@Param			userId		query		string							false	"ID пользователя (UUID)"
//	@Param			serviceName	query		string							false	"Название сервиса"
//	@Success		200			{object}	response.TotalPriceResponse		"Сумма всех подписок"
//	@Failure		400			{object}	response.BadRequestError		"Bad Request"
//	@Failure		500			{object}	response.InternalServerError	"Internal Server Error"
//	@Router			/subscriptions/total [get]
func (h *Handler) Handle(c *gin.Context) {
	var userID *uuid.UUID

	userIDStr := c.Query("userId")
	serviceName := c.Query("serviceName")

	if userIDStr != "" {
		parsed, err := uuid.Parse(userIDStr)
		if err != nil {
			h.logger.Warn("invalid userId",
				zap.String("userId", userIDStr),
				zap.Error(err),
			)
			c.JSON(http.StatusBadRequest, response.BadRequestError{Error: "Bad Request"})
			return
		}

		userID = &parsed
	}

	total, err := h.usecase.GetTotalPrice(c.Request.Context(), userID, pointer.ToStringOrNil(serviceName))
	if err != nil {
		h.logger.Error("h.usecase.GetTotalPrice",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, response.InternalServerError{Error: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, response.TotalPriceResponse{TotalPrice: pointer.GetInt64(total)})
}
