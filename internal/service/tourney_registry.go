package service

import (
	"context"
	"github.com/pejuang-awan/BE-Authentication/internal/shared"
)

type (
	TourneyRegistry interface {
		Create(ctx context.Context, req interface{}) error
		Get(ctx context.Context, req interface{}) error
		Update(ctx context.Context, req interface{}) error
		Delete(ctx context.Context, req interface{}) error
	}

	tourneyRegistryService struct {
		deps shared.Deps
	}
)

func (t *tourneyRegistryService) Create(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyRegistryService) Get(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyRegistryService) Update(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyRegistryService) Delete(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func NewTourneyRegistry(deps shared.Deps) (TourneyRegistry, error) {
	return &tourneyRegistryService{
		deps: deps,
	}, nil
}
