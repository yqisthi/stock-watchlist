package main

import (
	"stock-watchlist/application/controller"
	"stock-watchlist/application/service"
	"stock-watchlist/domain/repository"
	"stock-watchlist/infrastructure/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.RunMigrations()
	database.ConnectDB()

	app := fiber.New()

	userRepo := repository.ProvideUserRepository(database.DB)
	authService := service.ProvideAuthService(userRepo)
	authController := controller.ProvideAuthController(authService)

	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)

	app.Listen(":8080")
}