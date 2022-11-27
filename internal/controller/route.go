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

	CreateTournamentAPI       = "/tournament"
	GetTournamentByIDAPI      = "/tournament/:id"
	GetTournamentsAPI         = "/tournaments"
	GetTournamentsByGameIDAPI = "/tournaments/:gameID"
	GetGameByIDAPI            = "/game/:id"
	GetGamesAPI               = "/games"

	JoinTournamentAPI             = "/join"
	GetParticipantsByTourneyIDAPI = "/participants/:tourneyID"
	GetTournamentsByCaptainIDAPI  = "/tournaments"
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

	tourneyManagerRoutes := app.Group(PrefixTourneyManagerAPI)
	tourneyManagerRoutes.Use(customMiddleware.AuthMiddleware)
	{
		tourneyManagerRoutes.POST(CreateTournamentAPI, h.TourneyManager.CreateTourney)
		tourneyManagerRoutes.GET(GetTournamentByIDAPI, h.TourneyManager.GetTourneyById)
		tourneyManagerRoutes.GET(GetTournamentsAPI, h.TourneyManager.GetAllTourney)
		tourneyManagerRoutes.GET(GetTournamentsByGameIDAPI, h.TourneyManager.GetAllTourneyByGameID)
		tourneyManagerRoutes.GET(GetGameByIDAPI, h.TourneyManager.GetGameById)
		tourneyManagerRoutes.GET(GetGamesAPI, h.TourneyManager.GetAllGames)
	}

	tourneyRegistryRoutes := app.Group(PrefixTourneyRegistryAPI)
	tourneyRegistryRoutes.Use(customMiddleware.AuthMiddleware)
	{
		tourneyRegistryRoutes.POST(JoinTournamentAPI, h.TourneyRegistry.JoinTourney)
		tourneyRegistryRoutes.GET(GetParticipantsByTourneyIDAPI, h.TourneyRegistry.GetAllParticipantsByTourneyID)
		tourneyRegistryRoutes.GET(GetTournamentsByCaptainIDAPI, h.TourneyRegistry.GetAllTourneysByCaptainID)
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
