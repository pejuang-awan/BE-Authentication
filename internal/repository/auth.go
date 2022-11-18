package repository

import (
	"github.com/pejuang-awan/BE-Authentication/internal/entity"
	"gorm.io/gorm"
)

type (
	Auth interface {
		CreateUser(orm *gorm.DB, user *entity.User) (*entity.User, error)
		FindUserByUsername(orm *gorm.DB, username string) (*entity.User, error)
		FindUserByEmail(orm *gorm.DB, email string) (*entity.User, error)
	}

	authRepo struct{}
)

func (a *authRepo) CreateUser(orm *gorm.DB, user *entity.User) (*entity.User, error) {
	if err := orm.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authRepo) FindUserByUsername(orm *gorm.DB, username string) (*entity.User, error) {
	var user = &entity.User{}
	if err := orm.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authRepo) FindUserByEmail(orm *gorm.DB, email string) (*entity.User, error) {
	var user = &entity.User{}
	if err := orm.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func NewAuth() (Auth, error) {
	return &authRepo{}, nil
}
