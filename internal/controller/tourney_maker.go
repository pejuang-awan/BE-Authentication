package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pejuang-awan/BE-Authentication/internal/service"
	"github.com/pejuang-awan/BE-Authentication/internal/shared/dto"
	"go.uber.org/dig"
	"net/http"
)

type TourneyMaker struct {
	dig.In
	Service service.Holder
}

func (impl *TourneyMaker) Create(c echo.Context) error {
	var (
		response interface{}
	)

	// TODO: implement this
	// Get request body and parse it to struct
	// Add username or role value if needed

	bytes, statusCode, err := impl.Service.TourneyMaker.Create(c.Request())
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

func (impl *TourneyMaker) Get(c echo.Context) error {
	return nil
}

func (impl *TourneyMaker) Update(c echo.Context) error {
	return nil
}

func (impl *TourneyMaker) Delete(c echo.Context) error {
	return nil
}
