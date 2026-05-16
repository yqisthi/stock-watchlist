package router

import (
	"stock-watchlist/application/controller"

	"github.com/gofiber/fiber/v2"
)

func setupAuthRoutes(router fiber.Router, authController controller.AuthControllerInterface) {
	authGroup := router.Group("/auth")
	authGroup.Post("/register", authController.Register)
	authGroup.Post("/login", authController.Login)
}