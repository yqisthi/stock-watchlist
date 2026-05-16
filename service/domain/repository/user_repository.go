package repository

import (
	"context"
	"stock-watchlist/application/errs"
	"stock-watchlist/domain/model/dao"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *dao.User) error
	GetByEmail(ctx context.Context, email string) (*dao.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(ctx context.Context, user *dao.User) (err error) {
	err = ur.DB.WithContext(ctx).
		Create(user).
		Error

	if err != nil {
		err = errs.NewDatabaseError(err)
		return err
	}

	return
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (user *dao.User, err error) {
	err = ur.DB.WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		err = errs.NewDatabaseError(err)
		return nil, err
	}

	return user, nil
}