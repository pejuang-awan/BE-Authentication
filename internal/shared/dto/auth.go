package dto

import "errors"

const (
	CreateUserSuccess = "User created successfully"
	LoginSuccess      = "Login success"
)

var (
	ErrFindUserByUsernameFailed = errors.New("failed to find user by username")
	ErrFindUserByEmailFailed    = errors.New("failed to find user by email")
	ErrUsernameAlreadyExists    = errors.New("username already exists")
	ErrEmailAlreadyExists       = errors.New("email already exists")
	ErrCreateUserFailed         = errors.New("failed to create user")
	ErrUserNotFound             = errors.New("user not found")
	ErrWrongPassword            = errors.New("wrong password")
	ErrGenerateTokenFailed      = errors.New("failed to generate token")
)

type (
	SignUpRequest struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,password"`
		Role     string `json:"role" validate:"required"`
	}

	SignInRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required,password"`
	}

	AuthResponse struct {
		ID       uint   `json:"id" validate:"required"`
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Role     string `json:"role" validate:"required"`
		Token    string `json:"token" validate:"required"`
	}
)
