package integration

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gratefultolord/users-subscriptions/internal/api/handlers/subscription_list"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/database"
	integration_testing "github.com/gratefultolord/users-subscriptions/internal/infrastructure/testing/integration"
	"github.com/gratefultolord/users-subscriptions/internal/usecases/get_subscriptions_list"
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
	testDataPath := filepath.Join(wd, "..", "..", "..", "..", "..", "migrations", "insert_test_subscriptions.sql")

	err = database.RunMigrations(db, migrationPath, testDataPath)
	require.NoError(t, err)

	logger := zap.NewNop()

	storage := get_subscriptions_list.NewStorage(db)
	usecase := get_subscriptions_list.NewUsecase(storage)
	handler := subscription_list.NewHandler(logger, usecase)

	req, err := http.NewRequest("GET", "/subscriptions", nil)
	require.NoError(t, err)

	req.Header.Add("Content-Type", "application/json")

	router := gin.Default()
	router.GET("/subscriptions", handler.Handle)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
