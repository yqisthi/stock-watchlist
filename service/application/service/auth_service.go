package service

import (
	"context"
	"stock-watchlist/application/helper"
	"stock-watchlist/domain/model/dao"
	"stock-watchlist/domain/model/dto"
	"stock-watchlist/domain/model/message"
	"stock-watchlist/domain/repository"

	"go.opentelemetry.io/otel"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, registerRequest *dto.RegisterRequest) error
	Login(ctx context.Context, loginRequest *dto.LoginRequest) (*dto.LoginResponse, error)
}

type AuthService struct {
	UserRepo repository.UserRepositoryInterface
}

func ProvideAuthService(userRepo repository.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{UserRepo: userRepo}
}

var tracer = otel.Tracer("auth-service")

func (s *AuthService) Register(ctx context.Context, registerRequest *dto.RegisterRequest) error {
	ctx, span := tracer.Start(ctx, "AuthService.Register")
	defer span.End()

	hashedPassword, err := helper.HashPassword(registerRequest.Password)
	if err != nil {
		span.RecordError(err)
		return err
	}

	err = s.UserRepo.Create(ctx, &dao.User{
		Name:     registerRequest.Username,
		Email:    registerRequest.Email,
		Password: hashedPassword,
	})

	if err != nil {
		span.RecordError(err)
	}

	return err
}

func (s *AuthService) Login(ctx context.Context, loginRequest *dto.LoginRequest) (*dto.LoginResponse, error) {
	ctx, span := tracer.Start(ctx, "AuthService.Login")
	defer span.End()

	user, err := s.UserRepo.GetByEmail(ctx, loginRequest.Email)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	if !helper.CheckPassword(loginRequest.Password, user.Password) {
		err := message.ErrInvalidCredentials.GetError()
		span.RecordError(err)
		return nil, err
	}

	token, err := helper.GenerateJWT(user.ID)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return &dto.LoginResponse{Token: token}, nil
}