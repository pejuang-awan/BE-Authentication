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

func (impl *TourneyManager) Create(c echo.Context) error {
	var (
		req      = dto.CreateTournamentRequest{}
		response interface{}
	)

	if err := bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Error: err.Error(),
		})
	}

	req.Organizer = c.Get("username").(string)

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error(),
		})
	}

	bytes, statusCode, err := impl.Service.TourneyManager.Create(reqBytes)
	if err != nil {
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

	return c.JSON(http.StatusOK, dto.Response{
		Data: response,
	})
}

func (impl *TourneyManager) Get(c echo.Context) error {
	return nil
}

func (impl *TourneyManager) Update(c echo.Context) error {
	return nil
}

func (impl *TourneyManager) Delete(c echo.Context) error {
	return nil
}
