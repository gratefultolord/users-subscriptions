package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/gratefultolord/users-subscriptions/internal/api/handlers/subscription_update"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/database"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/domain"
	integration_testing "github.com/gratefultolord/users-subscriptions/internal/infrastructure/testing/integration"
	"github.com/gratefultolord/users-subscriptions/internal/usecases/update_subscription"
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

	// путь к миграциям от рабочей директории теста
	migrationPath := filepath.Join(wd, "..", "..", "..", "..", "..", "migrations", "init.sql")
	testDataPath := filepath.Join(wd, "..", "..", "..", "..", "..", "migrations", "insert_test_subscriptions.sql")

	err = database.RunMigrations(db, migrationPath, testDataPath)
	require.NoError(t, err)

	logger := zap.NewNop()

	storage := update_subscription.NewStorage(db)
	usecase := update_subscription.NewUsecase(storage)
	handler := subscription_update.NewHandler(logger, usecase)

	body, err := json.Marshal(domain.SubscriptionDTO{
		ServiceName: "Netflix",
		Price:       1500,
		UserID:      "ff88446a-1832-4b26-adec-60f5db2aed65",
		StartDate:   "01-2025",
		EndDate:     pointer.To("01-2026"),
	})
	require.NoError(t, err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("/subscriptions/%d/update", 2), bytes.NewReader(body))
	require.NoError(t, err)

	req.Header.Add("Content-Type", "application/json")

	router := gin.Default()
	router.PUT("/subscriptions/:subscriptionId/update", handler.Handle)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
