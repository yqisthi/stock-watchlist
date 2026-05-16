package router

import (
	"stock-watchlist/application/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, controllers *controller.Controllers) {
	setupAuthRoutes(router, controllers.Auth)
}