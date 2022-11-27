package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
	customMiddleware "github.com/pejuang-awan/BE-Authentication/internal/middleware"
	"github.com/pejuang-awan/BE-Authentication/internal/shared"
	"go.uber.org/dig"
	"unicode"
	"unicode/utf8"
)

const (
	PrefixAuthAPI            = "/api/auth"
	PrefixTourneyManagerAPI  = "/api/tourney-manager"
	PrefixTourneyRegistryAPI = "/api/tourney-registry"

	SignUpAPI = "/sign-up"
	SignInAPI = "/sign-in"
)

type CustomValidator struct {
	validator *validator.Validate
}

type Holder struct {
	dig.In
	Deps            shared.Deps
	Auth            Auth
	TourneyManager  TourneyManager
	TourneyRegistry TourneyRegistry
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (h *Holder) RegisterRoutes() {
	var app = h.Deps.Server

	newValidator := initValidator()
	app.Validator = &CustomValidator{validator: newValidator}

	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	authRoutes := app.Group(PrefixAuthAPI)
	{
		authRoutes.POST(SignUpAPI, h.Auth.SignUp)
		authRoutes.POST(SignInAPI, h.Auth.SignIn)
	}

	authenticatedRoutes := app.Group("")
	authenticatedRoutes.Use(customMiddleware.AuthMiddleware)
	{
		tourneyManagerRoutes := app.Group(PrefixTourneyManagerAPI)
		{
			tourneyManagerRoutes.POST("", h.TourneyManager.Create)
			tourneyManagerRoutes.GET("", h.TourneyManager.Get)
			tourneyManagerRoutes.PUT("", h.TourneyManager.Update)
			tourneyManagerRoutes.DELETE("", h.TourneyManager.Delete)
		}

		tourneyRegistryRoutes := app.Group(PrefixTourneyRegistryAPI)
		{
			tourneyRegistryRoutes.POST("", h.TourneyRegistry.Create)
			tourneyRegistryRoutes.GET("", h.TourneyRegistry.Get)
			tourneyRegistryRoutes.PUT("", h.TourneyRegistry.Update)
			tourneyRegistryRoutes.DELETE("", h.TourneyRegistry.Delete)
		}
	}
}

func initValidator() *validator.Validate {
	v := validator.New()

	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		var (
			hasNumber         = false
			hasLetter         = false
			hasSpecialChar    = false
			hasSuitableLength = false
		)

		password := fl.Field().String()

		if utf8.RuneCountInString(password) >= 8 && utf8.RuneCountInString(password) <= 30 {
			hasSuitableLength = true
		}

		for _, char := range password {
			switch {
			case unicode.IsNumber(char):
				hasNumber = true
			case unicode.IsLetter(char):
				hasLetter = true
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				hasSpecialChar = true
			default:
				return false
			}
		}

		return hasNumber && hasLetter && hasSpecialChar && hasSuitableLength
	})

	return v
}
