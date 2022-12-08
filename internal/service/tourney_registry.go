package service

import (
	"bytes"
	"github.com/pejuang-awan/BE-Authentication/internal/shared"
	"io"
	"io/ioutil"
	"net/http"
)

type (
	TourneyRegistry interface {
		JoinTourney(reqBytes []byte) ([]byte, int, error)
		GetParticipantsByTourneyID(tourneyID string) ([]byte, int, error)
		GetTourneysByCaptainID(captainID string) ([]byte, int, error)
	}

	tourneyRegistryService struct {
		deps shared.Deps
	}
)

func (t *tourneyRegistryService) JoinTourney(reqBytes []byte) ([]byte, int, error) {
	url := t.deps.Config.Services.TourneyRegistryURL + joinTourneyURL

	t.deps.Logger.Infof("Calling tourney registry service url %s", url)

	response, err := httpCall(http.MethodPost, url, bytes.NewReader(reqBytes))
	if err != nil || response.StatusCode != http.StatusOK {
		t.deps.Logger.Errorf("Error when send request to join in tourney registry service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Info("Success when send request to join in tourney registry service")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.deps.Logger.Errorf("Error when close body: %v", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.deps.Logger.Errorf("Error when read response body in tourney registry service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Infof("Response from tourney registry service: %v", string(body))

	return body, response.StatusCode, nil
}

func (t *tourneyRegistryService) GetParticipantsByTourneyID(tourneyID string) ([]byte, int, error) {
	url := t.deps.Config.Services.TourneyRegistryURL + getParticipantsURL + "/" + tourneyID

	t.deps.Logger.Infof("Calling tourney registry service url %s", url)

	response, err := httpCall(http.MethodGet, url, nil)
	if err != nil || response.StatusCode != http.StatusOK {
		t.deps.Logger.Errorf("Error when send request to get participants in tourney registry service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Info("Success when send request to get participants in tourney registry service")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.deps.Logger.Errorf("Error when close body: %v", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.deps.Logger.Errorf("Error when read response body in tourney registry service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Infof("Response from tourney registry service: %v", string(body))

	return body, response.StatusCode, nil
}

func (t *tourneyRegistryService) GetTourneysByCaptainID(captainID string) ([]byte, int, error) {
	url := t.deps.Config.Services.TourneyRegistryURL + getTourneysURL + "/" + captainID

	t.deps.Logger.Infof("Calling tourney registry service url %s", url)
	t.deps.Logger.Infof("our captain id is %s", captainID)

	response, err := httpCall(http.MethodGet, url, nil)
	if err != nil || response.StatusCode != http.StatusOK {
		t.deps.Logger.Errorf("Error when send request to get tourneys in tourney registry service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Info("Success when send request to get tourneys in tourney registry service")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.deps.Logger.Errorf("Error when close body: %v", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.deps.Logger.Errorf("Error when read response body in tourney registry service: %v", err)
		return nil, response.StatusCode, err
	}

	t.deps.Logger.Infof("Response from tourney registry service: %v", string(body))

	return body, response.StatusCode, nil
}

func NewTourneyRegistry(deps shared.Deps) (TourneyRegistry, error) {
	return &tourneyRegistryService{
		deps: deps,
	}, nil
}
