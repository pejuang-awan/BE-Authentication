package service

import (
	"bytes"
	"github.com/pejuang-awan/BE-Authentication/internal/shared"
	"io/ioutil"
	"net/http"
)

type (
	TourneyManager interface {
		Create(reqBytes []byte) ([]byte, int, error)
		Get(req *http.Request) ([]byte, int, error)
		Update(req *http.Request) ([]byte, int, error)
		Delete(req *http.Request) ([]byte, int, error)
	}

	tourneyManagerService struct {
		deps shared.Deps
	}
)

func (t *tourneyManagerService) Create(reqBytes []byte) ([]byte, int, error) {
	url := t.deps.Config.Services.TourneyManagerURL + createTourney

	response, err := httpCall(http.MethodPost, url, bytes.NewReader(reqBytes))
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

func (t *tourneyManagerService) Get(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyManagerService) Update(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tourneyManagerService) Delete(req *http.Request) ([]byte, int, error) {
	//TODO implement me
	panic("implement me")
}

func NewTourneyManager(deps shared.Deps) (TourneyManager, error) {
	return &tourneyManagerService{
		deps: deps,
	}, nil
}
