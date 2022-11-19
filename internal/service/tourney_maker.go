package service

import (
	"github.com/pejuang-awan/BE-Authentication/internal/shared"
	"io/ioutil"
	"net/http"
)

type (
	TourneyMaker interface {
		Create(req *http.Request) ([]byte, int, error)
		Get(req *http.Request) ([]byte, int, error)
		Update(req *http.Request) ([]byte, int, error)
		Delete(req *http.Request) ([]byte, int, error)
	}

	tourneyMakerService struct {
		deps shared.Deps
	}
)

func (t *tourneyMakerService) Create(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	response, err := httpCall(http.MethodPost, t.deps.Config.Services.TourneyMakerURL, req.Body)
	if err != nil || response.StatusCode != http.StatusOK {
		t.deps.Logger.Errorf("Error when send request to create in tourney maker service: %v", err)
		return nil, response.StatusCode, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.deps.Logger.Errorf("Error when read response body in tourney maker service: %v", err)
		return nil, response.StatusCode, err
	}

	return body, response.StatusCode, nil
}

func (t *tourneyMakerService) Get(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyMakerService) Update(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyMakerService) Delete(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	panic("implement me")
}

func NewTourneyMaker(deps shared.Deps) (TourneyMaker, error) {
	return &tourneyMakerService{
		deps: deps,
	}, nil
}
