package service

import (
	"github.com/pejuang-awan/BE-Authentication/internal/shared"
	"net/http"
)

type (
	TourneyRegistry interface {
		Create(req *http.Request) ([]byte, int, error)
		Get(req *http.Request) ([]byte, int, error)
		Update(req *http.Request) ([]byte, int, error)
		Delete(req *http.Request) ([]byte, int, error)
	}

	tourneyRegistryService struct {
		deps shared.Deps
	}
)

func (t *tourneyRegistryService) Create(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyRegistryService) Get(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyRegistryService) Update(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyRegistryService) Delete(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	panic("implement me")
}

func NewTourneyRegistry(deps shared.Deps) (TourneyRegistry, error) {
	return &tourneyRegistryService{
		deps: deps,
	}, nil
}
