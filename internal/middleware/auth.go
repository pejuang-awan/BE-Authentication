package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pejuang-awan/BE-Authentication/internal/service"
	"github.com/pejuang-awan/BE-Authentication/internal/shared/dto"
	"net/http"
)

const (
	BearerLength = 7
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if len(token) < BearerLength {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Error: dto.ErrUnauthorized.Error(),
			})
		}

		token = token[BearerLength:]
		claims := &service.Claims{}
		jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return service.JWTKey, nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Error: dto.ErrUnauthorized.Error(),
			})
		}

		if !jwtToken.Valid {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Error: dto.ErrUnauthorized.Error(),
			})
		}

		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		return next(c)
	}
}
