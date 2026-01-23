package dto

type AuthResponse struct {
	Token string   `json:"token"`
	User  UserData `json:"user"`
}

type UserData struct {
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}
