package service

import (
	"bytes"
	"github.com/pejuang-awan/BE-Authentication/internal/shared"
	"io"
	"io/ioutil"
	"net/http"
)

type (
	TourneyManager interface {
		CreateTourney(reqBytes []byte) ([]byte, int, error)
		GetTourneyById(id string) ([]byte, int, error)
		GetTourneys() ([]byte, int, error)
		GetTourneysByGameId(gameId string) ([]byte, int, error)
		GetGameById(gameId string) ([]byte, int, error)
		GetGames() ([]byte, int, error)
	}

	tourneyManagerService struct {
		deps shared.Deps
	}
)

func (t *tourneyManagerService) CreateTourney(reqBytes []byte) ([]byte, int, error) {
	url := t.deps.Config.Services.TourneyManagerURL + tourneyURL

	t.deps.Logger.Infof("Calling tourney manager service url %s", url)

	response, err := httpCall(http.MethodPost, url, bytes.NewReader(reqBytes))
	if err != nil || response.StatusCode != http.StatusOK {
		t.deps.Logger.Errorf("Error when send request to create in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Info("Success when send request to create in tourney manager service")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.deps.Logger.Errorf("Error when close body: %v", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.deps.Logger.Errorf("Error when read response body in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Infof("Response from tourney manager service: %v", string(body))

	return body, response.StatusCode, nil
}

func (t *tourneyManagerService) GetTourneyById(id string) ([]byte, int, error) {
	url := t.deps.Config.Services.TourneyManagerURL + tourneyURL + "/" + id

	t.deps.Logger.Infof("Calling tourney manager service url %s", url)

	response, err := httpCall(http.MethodGet, url, nil)
	if err != nil || response.StatusCode != http.StatusOK {
		t.deps.Logger.Errorf("Error when send request to get tourney by id in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Info("Success when send request to get tourney by id in tourney manager service")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.deps.Logger.Errorf("Error when close body: %v", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.deps.Logger.Errorf("Error when read response body in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Infof("Response from tourney manager service: %v", string(body))

	return body, response.StatusCode, nil
}

func (t *tourneyManagerService) GetTourneys() ([]byte, int, error) {
	url := t.deps.Config.Services.TourneyManagerURL + tourneysURL

	t.deps.Logger.Infof("Calling tourney manager service url %s", url)

	response, err := httpCall(http.MethodGet, url, nil)
	if err != nil || response.StatusCode != http.StatusOK {
		t.deps.Logger.Errorf("Error when send request to get all tourney in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Info("Success when send request to get all tourney in tourney manager service")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.deps.Logger.Errorf("Error when close body: %v", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.deps.Logger.Errorf("Error when read response body in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Infof("Response from tourney manager service: %v", string(body))

	return body, response.StatusCode, nil
}

func (t *tourneyManagerService) GetTourneysByGameId(gameId string) ([]byte, int, error) {
	url := t.deps.Config.Services.TourneyManagerURL + tourneysURL + "/" + gameId

	t.deps.Logger.Infof("Calling tourney manager service url %s", url)

	response, err := httpCall(http.MethodGet, url, nil)
	if err != nil || response.StatusCode != http.StatusOK {
		t.deps.Logger.Errorf("Error when send request to get tourney by game id in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Info("Success when send request to get tourney by game id in tourney manager service")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.deps.Logger.Errorf("Error when close body: %v", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.deps.Logger.Errorf("Error when read response body in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Infof("Response from tourney manager service: %v", string(body))

	return body, response.StatusCode, nil
}

func (t *tourneyManagerService) GetGameById(gameId string) ([]byte, int, error) {
	url := t.deps.Config.Services.TourneyManagerURL + gameURL + "/" + gameId

	t.deps.Logger.Infof("Calling tourney manager service url %s", url)

	response, err := httpCall(http.MethodGet, url, nil)
	if err != nil || response.StatusCode != http.StatusOK {
		t.deps.Logger.Errorf("Error when send request to get game by id in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Info("Success when send request to get game by id in tourney manager service")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.deps.Logger.Errorf("Error when close body: %v", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.deps.Logger.Errorf("Error when read response body in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Infof("Response from tourney manager service: %v", string(body))

	return body, response.StatusCode, nil
}

func (t *tourneyManagerService) GetGames() ([]byte, int, error) {
	url := t.deps.Config.Services.TourneyManagerURL + gamesURL

	t.deps.Logger.Infof("Calling tourney manager service url %s", url)

	response, err := httpCall(http.MethodGet, url, nil)
	if err != nil || response.StatusCode != http.StatusOK {
		t.deps.Logger.Errorf("Error when send request to get all games in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Info("Success when send request to get all games in tourney manager service")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.deps.Logger.Errorf("Error when close body: %v", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.deps.Logger.Errorf("Error when read response body in tourney manager service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Infof("Response from tourney manager service: %v", string(body))

	return body, response.StatusCode, nil
}

func NewTourneyManager(deps shared.Deps) (TourneyManager, error) {
	return &tourneyManagerService{
		deps: deps,
	}, nil
}
