package service

import (
	"context"
	"github.com/pejuang-awan/BE-Authentication/internal/entity"
	"github.com/pejuang-awan/BE-Authentication/internal/repository"
	"github.com/pejuang-awan/BE-Authentication/internal/shared"
	"github.com/pejuang-awan/BE-Authentication/internal/shared/dto"
)

type (
	Auth interface {
		CreateUser(ctx context.Context, req *dto.TestRequest) (*dto.TestResponse, error)
	}

	authService struct {
		deps shared.Deps
		repo repository.Auth
	}
)

func (a *authService) CreateUser(ctx context.Context, req *dto.TestRequest) (*dto.TestResponse, error) {
	var orm = a.deps.Database.WithContext(ctx)
	userCreated, err := a.repo.CreateUser(orm, &entity.User{Username: req.Name, Email: "test@mail.com", Password: "password", Role: "user"})

	if err != nil {
		a.deps.Logger.Errorf("Failed to create user: %v", err)
		return nil, err
	}

	return &dto.TestResponse{
		ID:   userCreated.ID,
		Name: userCreated.Username,
	}, nil

}

func NewAuth(deps shared.Deps, repo repository.Auth) (Auth, error) {
	return &authService{deps, repo}, nil
}
