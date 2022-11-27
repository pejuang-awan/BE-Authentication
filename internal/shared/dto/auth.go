package dto

import "errors"

var (
	ErrFindUserByUsernameFailed = errors.New("failed to find user by username")
	ErrUsernameAlreadyExists    = errors.New("username already exists")
	ErrCreateUserFailed         = errors.New("failed to create user")
	ErrUserNotFound             = errors.New("user not found")
	ErrWrongPassword            = errors.New("wrong password")
	ErrGenerateTokenFailed      = errors.New("failed to generate token")
	ErrUnauthorized             = errors.New("unauthorized")
)

type (
	SignUpRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required,password"`
		Role     uint8  `json:"role" validate:"required"`
		GameType uint8  `json:"gameType" validate:"required"`
	}

	SignInRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required,password"`
	}

	AuthResponse struct {
		ID       uint   `json:"id" validate:"required"`
		Username string `json:"username" validate:"required"`
		Role     uint8  `json:"role" validate:"required"`
		GameType uint8  `json:"gameType" validate:"required"`
		Token    string `json:"token" validate:"required"`
	}
)
