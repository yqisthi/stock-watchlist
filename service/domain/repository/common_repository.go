package repository

import (
	"stock-watchlist/tracing"

	"gorm.io/gorm"
)

type Repositories struct {
	User UserRepositoryInterface
}

func ProvideRepositories(db *gorm.DB) *Repositories {
	userRepository := ProvideUserRepository(db)
	userRepository = NewUserRepositoryOtel(tracing.TracerRepository, userRepository)

	return &Repositories{
		User: userRepository,
	}
}