package service

import (
	"context"
	"errors"
	"stock-watchlist/application/helper"
	"stock-watchlist/domain/model/dao"
	"stock-watchlist/domain/model/dto"
	"stock-watchlist/domain/repository"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, registerRequest dto.RegisterRequest) error
	Login(ctx context.Context, loginRequest dto.LoginRequest) (*dto.LoginResponse, error)
}

type AuthService struct {
	UserRepo repository.UserRepositoryInterface
}

func ProvideAuthService(userRepo repository.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Register(ctx context.Context, registerRequest dto.RegisterRequest) error {
	return s.UserRepo.Create(ctx, &dao.User{
		Name:     registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	})
}

func (s *AuthService) Login(ctx context.Context, loginRequest dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.UserRepo.GetByEmail(ctx, loginRequest.Email)
	if err != nil {
		return nil, err
	}

	if !helper.CheckPassword(loginRequest.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := helper.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: token}, nil
}