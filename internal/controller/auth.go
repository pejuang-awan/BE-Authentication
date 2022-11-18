package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pejuang-awan/BE-Authentication/internal/service"
	"github.com/pejuang-awan/BE-Authentication/internal/shared/dto"
	"go.uber.org/dig"
	"net/http"
)

type Auth struct {
	dig.In
	Service service.Holder
}

func (impl *Auth) Post(c echo.Context) error {
	var req = dto.TestRequest{}

	if err := bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	res, err := impl.Service.Auth.CreateUser(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  dto.StatusError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Status:  dto.StatusSuccess,
		Message: dto.CreateUserSuccess,
		Data:    res,
	})
}
