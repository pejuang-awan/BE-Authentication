package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pejuang-awan/BE-Authentication/internal/service"
	"go.uber.org/dig"
)

type TourneyRegistry struct {
	dig.In
	Service service.Holder
}

func (impl *TourneyRegistry) Create(c echo.Context) error {
	return nil
}

func (impl *TourneyRegistry) Get(c echo.Context) error {
	return nil
}

func (impl *TourneyRegistry) Update(c echo.Context) error {
	return nil
}

func (impl *TourneyRegistry) Delete(c echo.Context) error {
	return nil
}
