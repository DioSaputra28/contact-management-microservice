package dto

// LoginRequest represents login input data
type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

// RegisterRequest represents registration input data
type RegisterRequest struct {
	Name     string `validate:"required,min=1,max=100"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}
