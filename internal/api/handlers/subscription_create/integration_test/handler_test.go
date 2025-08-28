package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/gratefultolord/users-subscriptions/internal/api/handlers/subscription_create"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/database"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
	integration_testing "github.com/gratefultolord/users-subscriptions/internal/infrastructure/testing/integration"
	"github.com/gratefultolord/users-subscriptions/internal/usecases/create_subscription"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestHandler_Handle(t *testing.T) {
	pg := integration_testing.NewPostgres(t)
	defer pg.CleanUp()

	db := sqlx.NewDb(pg.DB, "postgres")

	wd, err := os.Getwd()
	require.NoError(t, err)

	// путь к миграции от рабочей директории теста
	migrationPath := filepath.Join(wd, "..", "..", "..", "..", "..", "migrations", "init.sql")

	err = database.RunMigrations(db, migrationPath)
	require.NoError(t, err)

	logger := zap.NewNop()

	storage := create_subscription.NewStorage(db)
	usecase := create_subscription.NewUsecase(storage)
	handler := subscription_create.NewHandler(logger, usecase)

	body, err := json.Marshal(domain.SubscriptionDTO{
		ServiceName: "Yandex Plus",
		Price:       400,
		UserID:      "ff88446a-1832-4b26-adec-60f5db2aed65",
		StartDate:   "06-2025",
		EndDate:     pointer.To("10-2025"),
	})
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/subscriptions/create", bytes.NewReader(body))
	require.NoError(t, err)

	req.Header.Add("Content-Type", "application/json")

	router := gin.Default()
	router.POST("/subscriptions/create", handler.Handle)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
