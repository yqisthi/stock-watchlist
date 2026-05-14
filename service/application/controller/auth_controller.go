package controller

import (
	"stock-watchlist/application/service"
	"stock-watchlist/domain/model/dto"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerInterface interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

type AuthController struct {
	AuthService service.AuthServiceInterface
}

func ProvideAuthController(authService service.AuthServiceInterface) AuthControllerInterface {
	return &AuthController{AuthService: authService}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var registerRequest dto.RegisterRequest

	if err := ctx.BodyParser(&registerRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	err := c.AuthService.Register(ctx.Context(), registerRequest)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "register success",
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var loginRequest dto.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	token, err := c.AuthService.Login(ctx.Context(), loginRequest)

	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(token)
}