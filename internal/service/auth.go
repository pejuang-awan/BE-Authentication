package service

import (
	"context"
	"github.com/pejuang-awan/BE-Authentication/internal/entity"
	"github.com/pejuang-awan/BE-Authentication/internal/repository"
	"github.com/pejuang-awan/BE-Authentication/internal/shared"
	"github.com/pejuang-awan/BE-Authentication/internal/shared/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type (
	Auth interface {
		SignUp(ctx context.Context, req *dto.SignUpRequest) (*dto.AuthResponse, int, error)
		SignIn(ctx context.Context, req *dto.SignInRequest) (*dto.AuthResponse, int, error)
	}

	authService struct {
		deps shared.Deps
		repo repository.Auth
	}
)

func (a *authService) SignUp(ctx context.Context, req *dto.SignUpRequest) (*dto.AuthResponse, int, error) {
	var orm = a.deps.Database.WithContext(ctx)

	oldUsername, err := a.repo.FindUserByUsername(orm, req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		a.deps.Logger.Errorf("Error while finding user by username: %v", err)
		return nil, http.StatusInternalServerError, dto.ErrFindUserByUsernameFailed
	}

	if oldUsername != nil {
		a.deps.Logger.Error("Username already exists")
		return nil, http.StatusConflict, dto.ErrUsernameAlreadyExists
	}

	oldEmail, err := a.repo.FindUserByEmail(orm, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		a.deps.Logger.Errorf("Error while finding user by email: %v", err)
		return nil, http.StatusInternalServerError, dto.ErrFindUserByEmailFailed
	}

	if oldEmail != nil {
		a.deps.Logger.Error("Email already exists")
		return nil, http.StatusConflict, dto.ErrEmailAlreadyExists
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	hashedPasswordString := string(hashedPassword)

	userCreated, err := a.repo.CreateUser(orm, &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPasswordString,
		Role:     req.Role,
	})

	if err != nil {
		a.deps.Logger.Errorf("Error while creating user: %v", err)
		return nil, http.StatusInternalServerError, dto.ErrCreateUserFailed
	}

	return &dto.AuthResponse{
		ID:       userCreated.ID,
		Username: userCreated.Username,
		Email:    userCreated.Email,
		Role:     userCreated.Role,
		Token:    "token",
	}, http.StatusOK, nil
}

func (a *authService) SignIn(ctx context.Context, req *dto.SignInRequest) (*dto.AuthResponse, int, error) {
	var orm = a.deps.Database.WithContext(ctx)

	user, err := a.repo.FindUserByUsername(orm, req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			a.deps.Logger.Error("User not found")
			return nil, http.StatusNotFound, dto.ErrUserNotFound
		}

		a.deps.Logger.Errorf("Error while finding user by username: %v", err)
		return nil, http.StatusInternalServerError, dto.ErrFindUserByUsernameFailed
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		a.deps.Logger.Error("Wrong password")
		return nil, http.StatusUnauthorized, dto.ErrWrongPassword
	}

	return &dto.AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Token:    "token",
	}, http.StatusOK, nil
}

func NewAuth(deps shared.Deps, repo repository.Auth) (Auth, error) {
	return &authService{deps, repo}, nil
}
