package service

import (
	"stock-watchlist/domain/repository"
	"stock-watchlist/tracing"
)

type Services struct {
	Auth AuthServiceInterface
}

func ProvideServices(repositories *repository.Repositories) *Services {
	authService := ProvideAuthService(repositories.User)
	authService = NewAuthServiceOtel(tracing.TracerService, authService)

	return &Services{
		Auth: authService,
	}
}