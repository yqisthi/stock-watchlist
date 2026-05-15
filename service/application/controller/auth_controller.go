package controller

import (
	"stock-watchlist/application/helper"
	"stock-watchlist/application/service"
	"stock-watchlist/domain/model/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthControllerInterface interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

type AuthController struct {
	AuthService service.AuthServiceInterface
	validator *validator.Validate
}

func ProvideAuthController(authService service.AuthServiceInterface, validator *validator.Validate) AuthControllerInterface {
	return &AuthController{AuthService: authService, validator: validator}
}

func (ac *AuthController) Register(ctx *fiber.Ctx) error {
	registerRequest, err := helper.ParseRequest[dto.RegisterRequest](ctx)
		if err != nil {
			return helper.ClientErrorJSONResponse(ctx, err)
	}

	if err = helper.ValidateRequest(registerRequest, ac.validator); err != nil {
		return helper.ClientErrorJSONResponse(ctx, err)
	}

	err = ac.AuthService.Register(ctx.Context(), registerRequest)

	if err != nil {
		return helper.ClientErrorJSONResponse(ctx, err)
	}

	return helper.ClientSuccessJSONResponse(ctx	, nil)
}

func (ac *AuthController) Login(ctx *fiber.Ctx) error {
	loginRequest, err := helper.ParseRequest[dto.LoginRequest](ctx)

	if err = helper.ValidateRequest(loginRequest, ac.validator); err != nil {
		return helper.ClientErrorJSONResponse(ctx, err)
	}

	token, err := ac.AuthService.Login(ctx.Context(), loginRequest)

	if err != nil {
		return helper.ClientErrorJSONResponse(ctx, err)
	}

	return helper.ClientSuccessJSONResponse(ctx, token)
}