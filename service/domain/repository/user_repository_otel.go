package repository

import (
	"context"
	"stock-watchlist/domain/model/dao"

	"go.opentelemetry.io/otel/trace"
)

type UserRepositoryOtelDecoratorInterface interface {
	UserRepositoryInterface
}

type userRepositoryOtel struct {
	tracerRepository trace.Tracer
	userRepository UserRepositoryInterface
}

func NewUserRepositoryOtel(repositoryTracer trace.Tracer, userRepository UserRepositoryInterface) UserRepositoryOtelDecoratorInterface {
	return &userRepositoryOtel{
		tracerRepository: repositoryTracer,
		userRepository:   userRepository,
	}
}

func (uro *userRepositoryOtel) Create(ctx context.Context, user *dao.User) error {
	ctx, span := uro.tracerRepository.Start(ctx, "UserRepository.Create")
	defer span.End()

	err := uro.userRepository.Create(ctx, user)
	if err != nil {
		span.RecordError(err)
	}

	return err
}

func (uro *userRepositoryOtel) GetByEmail(ctx context.Context, email string) (*dao.User, error) {
	ctx, span := uro.tracerRepository.Start(ctx, "UserRepository.GetByEmail")
	defer span.End()

	user, err := uro.userRepository.GetByEmail(ctx, email)
	if err != nil {
		span.RecordError(err)
	}

	return user, err
}