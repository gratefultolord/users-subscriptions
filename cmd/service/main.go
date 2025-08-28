package main

import (
	"log"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/gratefultolord/users-subscriptions/docs"
	"github.com/gratefultolord/users-subscriptions/internal/api/handlers/subscription_create"
	"github.com/gratefultolord/users-subscriptions/internal/api/handlers/subscription_delete"
	"github.com/gratefultolord/users-subscriptions/internal/api/handlers/subscription_list"
	"github.com/gratefultolord/users-subscriptions/internal/api/handlers/subscription_read"
	"github.com/gratefultolord/users-subscriptions/internal/api/handlers/subscription_update"
	"github.com/gratefultolord/users-subscriptions/internal/api/handlers/total_price_get"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/config"
	"github.com/gratefultolord/users-subscriptions/internal/infrastructure/database"
	"github.com/gratefultolord/users-subscriptions/internal/usecases/create_subscription"
	"github.com/gratefultolord/users-subscriptions/internal/usecases/delete_subscription"
	"github.com/gratefultolord/users-subscriptions/internal/usecases/get_subscription"
	"github.com/gratefultolord/users-subscriptions/internal/usecases/get_subscriptions_list"
	"github.com/gratefultolord/users-subscriptions/internal/usecases/get_total_price"
	"github.com/gratefultolord/users-subscriptions/internal/usecases/update_subscription"
	swagger_files "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

//	@title		Сервис работы с информацией о подписках
//	@version	1.0

//	@contact.name	Shukurillo Karimov
//	@contact.email	gratefultolord@gmail.com

//	@host	localhost:8080

func main() {
	loggerConfig := zap.NewProductionConfig()

	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "debug":
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatalf("zap.NewProduction: %v", err)
	}
	defer logger.Sync()

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("config.Load: %v", zap.Error(err))
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		logger.Fatal("database.NewPostgresDB: %v", zap.Error(err))
	}
	defer db.Close()

	// При необходимости убрать второй файл миграций (тестовые данные)
	err = database.RunMigrations(db, "migrations/init.sql", "migrations/insert_test_subscriptions.sql")
	if err != nil {
		logger.Fatal("database.RunMigrations: %v", zap.Error(err))
	}

	createSubscriptionStorage := create_subscription.NewStorage(db)
	deleteSubscriptionStorage := delete_subscription.NewStorage(db)
	getSubscriptionsListStorage := get_subscriptions_list.NewStorage(db)
	getSubscriptionStorage := get_subscription.NewStorage(db)
	updateSubscriptionStorage := update_subscription.NewStorage(db)
	getTotalPriceStorage := get_total_price.NewStorage(db)

	createSubscriptionUsecase := create_subscription.NewUsecase(createSubscriptionStorage)
	deleteSubscriptionUsecase := delete_subscription.NewUsecase(deleteSubscriptionStorage)
	getSubscriptionsListUsecase := get_subscriptions_list.NewUsecase(getSubscriptionsListStorage)
	getSubscriptionUsecase := get_subscription.NewUsecase(getSubscriptionStorage)
	updateSubscriptionUsecase := update_subscription.NewUsecase(updateSubscriptionStorage)
	getTotalPriceUsecase := get_total_price.NewUsecase(getTotalPriceStorage)

	createSubscriptionHandler := subscription_create.NewHandler(logger, createSubscriptionUsecase)
	deleteSubscriptionHandler := subscription_delete.NewHandler(logger, deleteSubscriptionUsecase)
	getSubscriptionsListHandler := subscription_list.NewHandler(logger, getSubscriptionsListUsecase)
	getSubscriptionHandler := subscription_read.NewHandler(logger, getSubscriptionUsecase)
	updateSubscriptionHandler := subscription_update.NewHandler(logger, updateSubscriptionUsecase)
	getTotalPriceHandler := total_price_get.NewHandler(logger, getTotalPriceUsecase)

	router := gin.New()

	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	router.GET("/swagger/*any", gin_swagger.WrapHandler(swagger_files.Handler))
	router.POST("/subscriptions/create", createSubscriptionHandler.Handle)
	router.DELETE("/subscriptions/:subscriptionId/delete", deleteSubscriptionHandler.Handle)
	router.GET("/subscriptions", getSubscriptionsListHandler.Handle)
	router.GET("/subscriptions/:subscriptionId", getSubscriptionHandler.Handle)
	router.PUT("/subscriptions/:subscriptionId/update", updateSubscriptionHandler.Handle)
	router.GET("/subscriptions/total", getTotalPriceHandler.Handle)

	err = router.Run(cfg.HTTPAddress)
	if err != nil {
		logger.Fatal("router.Run: %v", zap.Error(err))
	}

	logger.Info("Server is ready for connections", zap.String("addr", cfg.HTTPAddress))
}
