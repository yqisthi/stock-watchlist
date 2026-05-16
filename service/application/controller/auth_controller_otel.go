package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

type AuthControllerOtelDecoratorInterface interface {
	AuthControllerInterface
}

type authControllerOtel struct {
	tracerController trace.Tracer
	authController AuthControllerInterface
}

func NewAuthControllerOtel(tracerController trace.Tracer, authController AuthControllerInterface) AuthControllerOtelDecoratorInterface {
	return &authControllerOtel{
		tracerController: tracerController,
		authController:   authController,
	}
}

func (aco *authControllerOtel) Register(ctx *fiber.Ctx) (err error) {
	c, span := aco.tracerController.Start(ctx.UserContext(), "AuthController.Register")
	ctx.SetUserContext(c)
	defer span.End()

	err = aco.authController.Register(ctx)
	if err != nil {
		span.RecordError(err)
	}

	return 
}

func (aco *authControllerOtel) Login(ctx *fiber.Ctx) (err error) {
	c, span := aco.tracerController.Start(ctx.UserContext(), "AuthController.Login")
	ctx.SetUserContext(c)
	defer span.End()

	err = aco.authController.Login(ctx)
	if err != nil {
		span.RecordError(err)
	}

	return 
}