package controller

import (
	"stock-watchlist/application/service"
	"stock-watchlist/tracing"

	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	Auth AuthControllerInterface
}

func ProvideControllers(services *service.Services, structValidator *validator.Validate) *Controllers {
	authController := ProvideAuthController(services.Auth, structValidator)
	authController = NewAuthControllerOtel(tracing.TracerController, authController)
	
	return &Controllers{
		Auth: authController,
	}
}