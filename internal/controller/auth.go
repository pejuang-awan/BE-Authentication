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

func (impl *Auth) SignUp(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = dto.SignUpRequest{}
	)

	if err := bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message: err.Error(),
			Status:  dto.StatusError,
		})
	}

	res, statusCode, err := impl.Service.Auth.SignUp(ctx, &req)
	if err != nil {
		return c.JSON(statusCode, dto.Response{
			Message: err.Error(),
			Status:  dto.StatusError,
		})
	}

	if err := c.Validate(res); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message: err.Error(),
			Status:  dto.StatusError,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: dto.CreateUserSuccess,
		Status:  dto.StatusSuccess,
		Data:    res,
	})
}

func (impl *Auth) SignIn(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = dto.SignInRequest{}
	)

	if err := bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message: err.Error(),
			Status:  dto.StatusError,
		})
	}

	res, statusCode, err := impl.Service.Auth.SignIn(ctx, &req)
	if err != nil {
		return c.JSON(statusCode, dto.Response{
			Message: err.Error(),
			Status:  dto.StatusError,
		})
	}

	if err := c.Validate(res); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message: err.Error(),
			Status:  dto.StatusError,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: dto.LoginSuccess,
		Status:  dto.StatusSuccess,
		Data:    res,
	})
}
