package service

import (
	"context"
	"stock-watchlist/domain/model/dto"

	"go.opentelemetry.io/otel/trace"
)

type AuthServiceOtelDecoratorInterface interface {
	AuthServiceInterface
}

type AuthServiceOtel struct {
	tracerService trace.Tracer
	authService AuthServiceInterface
}

func NewAuthServiceOtel(tracerService trace.Tracer, authService AuthServiceInterface) AuthServiceOtelDecoratorInterface {
	return &AuthServiceOtel{
		tracerService: tracerService,
		authService:   authService,
	}
}

func (aso *AuthServiceOtel) Register(ctx context.Context, registerRequest *dto.RegisterRequest) error {
	ctx, span := aso.tracerService.Start(ctx, "AuthService.Register")
	defer span.End()

	err := aso.authService.Register(ctx, registerRequest)
	if err != nil {
		span.RecordError(err)
	}

	return err
}

func (aso *AuthServiceOtel) Login(ctx context.Context, loginRequest *dto.LoginRequest) (*dto.LoginResponse, error) {
	ctx, span := aso.tracerService.Start(ctx, "AuthService.Login")
	defer span.End()

	loginResponse, err := aso.authService.Login(ctx, loginRequest)
	if err != nil {
		span.RecordError(err)
	}

	return loginResponse, err
}