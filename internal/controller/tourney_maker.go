package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pejuang-awan/BE-Authentication/internal/service"
	"go.uber.org/dig"
)

type TourneyMaker struct {
	dig.In
	Service service.Holder
}

func (impl *TourneyMaker) Create(c echo.Context) error {
	return nil
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
