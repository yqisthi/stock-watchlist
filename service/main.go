package main

import (
	"context"
	"stock-watchlist/application/controller"
	"stock-watchlist/application/router"
	"stock-watchlist/application/service"
	"stock-watchlist/domain/repository"
	"stock-watchlist/infrastructure/database"
	"stock-watchlist/infrastructure/logger"
	"stock-watchlist/infrastructure/telemetry"
	"stock-watchlist/tracing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.InitLogger()

	database.RunMigrations()
	database.ConnectDB()

	shutdown := telemetry.InitTracer()
	defer shutdown(context.Background())
	tracing.Init()

	app := fiber.New()
	app.Use(otelfiber.Middleware())

	validator := validator.New()
	repositories := repository.ProvideRepositories(database.DB)
	services := service.ProvideServices(repositories)
	controllers := controller.ProvideControllers(services, validator)

	router.SetupRoutes(app, controllers)

	app.Listen(":8080")
}