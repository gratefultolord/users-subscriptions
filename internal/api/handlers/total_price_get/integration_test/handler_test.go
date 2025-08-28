package integration

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gratefultolord/users-subscriptions/internal/api/handlers/total_price_get"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/database"
	integration_testing "github.com/gratefultolord/users-subscriptions/internal/infrastructure/testing/integration"
	"github.com/gratefultolord/users-subscriptions/internal/usecases/get_total_price"
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

	storage := get_total_price.NewStorage(db)
	usecase := get_total_price.NewUsecase(storage)
	handler := total_price_get.NewHandler(logger, usecase)

	params := url.Values{}
	params.Add("userId", "22222222-2222-2222-2222-222222222222")
	params.Add("serviceName", "Netflix")

	req, err := http.NewRequest("GET", "/subscriptions/total?"+params.Encode(), nil)
	require.NoError(t, err)

	req.Header.Add("Content-Type", "application/json")

	router := gin.Default()
	router.GET("/subscriptions/total", handler.Handle)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
