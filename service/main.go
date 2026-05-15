package main

import (
	"context"
	"stock-watchlist/application/controller"
	"stock-watchlist/application/service"
	"stock-watchlist/domain/repository"
	"stock-watchlist/infrastructure/database"
	"stock-watchlist/infrastructure/logger"
	"stock-watchlist/infrastructure/telemetry"

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

	app := fiber.New()

	app.Use(otelfiber.Middleware())

	validator := validator.New()
	userRepo := repository.ProvideUserRepository(database.DB)
	authService := service.ProvideAuthService(userRepo)
	authController := controller.ProvideAuthController(authService, validator)

	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)

	app.Listen(":8080")
}