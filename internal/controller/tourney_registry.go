package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pejuang-awan/BE-Authentication/internal/service"
	"github.com/pejuang-awan/BE-Authentication/internal/shared/dto"
	"go.uber.org/dig"
	"net/http"
)

type TourneyRegistry struct {
	dig.In
	Service service.Holder
}

func (impl *TourneyRegistry) JoinTourney(c echo.Context) error {
	var (
		req      = dto.JoinTournamentRequest{}
		username = c.Get("username").(string)
		response interface{}
	)

	if err := bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Error: err.Error(),
		})
	}

	req.CaptainID = username

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error(),
		})
	}

	bytes, statusCode, err := impl.Service.TourneyRegistry.JoinTourney(reqBytes)
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

func (impl *TourneyRegistry) GetAllParticipantsByTourneyID(c echo.Context) error {
	var (
		response  interface{}
		tourneyID = c.Param("tourneyID")
	)

	bytes, statusCode, err := impl.Service.TourneyRegistry.GetParticipantsByTourneyID(tourneyID)
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

func (impl *TourneyRegistry) GetAllTourneysByCaptainID(c echo.Context) error {
	var (
		response  interface{}
		captainID = c.Get("username").(string)
	)

	bytes, statusCode, err := impl.Service.TourneyRegistry.GetTourneysByCaptainID(captainID)
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
