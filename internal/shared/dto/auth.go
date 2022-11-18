package dto

const (
	CreateUserSuccess = "User created successfully"
)

type (
	TestRequest struct {
		Name string `json:"name" validate:"required"`
	}

	TestResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
)
