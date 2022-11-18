package service

import (
	"context"
	"github.com/pejuang-awan/BE-Authentication/internal/shared"
)

type (
	TourneyMaker interface {
		Create(ctx context.Context, req interface{}) error
		Get(ctx context.Context, req interface{}) error
		Update(ctx context.Context, req interface{}) error
		Delete(ctx context.Context, req interface{}) error
	}

	tourneyMakerService struct {
		deps shared.Deps
	}
)

func (t *tourneyMakerService) Create(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyMakerService) Get(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyMakerService) Update(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyMakerService) Delete(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func NewTourneyMaker(deps shared.Deps) (TourneyMaker, error) {
	return &tourneyMakerService{
		deps: deps,
	}, nil
}
