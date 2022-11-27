package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pejuang-awan/BE-Authentication/internal/service"
	"github.com/pejuang-awan/BE-Authentication/internal/shared/dto"
	"go.uber.org/dig"
	"net/http"
)

type TourneyManager struct {
	dig.In
	Service service.Holder
}

func (impl *TourneyManager) CreateTourney(c echo.Context) error {
	var (
		req      = dto.CreateTournamentRequest{}
		username = c.Get("username").(string)
		response interface{}
	)

	if err := bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Error: err.Error(),
		})
	}

	req.Organizer = username

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error(),
		})
	}

	bytes, statusCode, err := impl.Service.TourneyManager.CreateTourney(reqBytes)
	if err != nil || statusCode != http.StatusOK {
		return c.JSON(statusCode, dto.Response{
			Error: err.Error(),
		})
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (impl *TourneyManager) GetTourneyById(c echo.Context) error {
	var (
		response  interface{}
		tourneyID = c.Param("id")
	)

	bytes, statusCode, err := impl.Service.TourneyManager.GetTourneyById(tourneyID)
	if err != nil || statusCode != http.StatusOK {
		return c.JSON(statusCode, dto.Response{
			Error: err.Error(),
		})
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (impl *TourneyManager) GetAllTourney(c echo.Context) error {
	var (
		response interface{}
	)

	bytes, statusCode, err := impl.Service.TourneyManager.GetTourneys()
	if err != nil || statusCode != http.StatusOK {
		return c.JSON(statusCode, dto.Response{
			Error: err.Error(),
		})
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (impl *TourneyManager) GetAllTourneyByGameID(c echo.Context) error {
	var (
		response interface{}
		gameID   = c.Param("gameId")
	)

	bytes, statusCode, err := impl.Service.TourneyManager.GetTourneysByGameId(gameID)
	if err != nil || statusCode != http.StatusOK {
		return c.JSON(statusCode, dto.Response{
			Error: err.Error(),
		})
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (impl *TourneyManager) GetGameById(c echo.Context) error {
	var (
		response interface{}
		gameID   = c.Param("id")
	)

	bytes, statusCode, err := impl.Service.TourneyManager.GetGameById(gameID)
	if err != nil || statusCode != http.StatusOK {
		return c.JSON(statusCode, dto.Response{
			Error: err.Error(),
		})
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (impl *TourneyManager) GetAllGames(c echo.Context) error {
	var (
		response interface{}
	)

	bytes, statusCode, err := impl.Service.TourneyManager.GetGames()
	if err != nil || statusCode != http.StatusOK {
		return c.JSON(statusCode, dto.Response{
			Error: err.Error(),
		})
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}
