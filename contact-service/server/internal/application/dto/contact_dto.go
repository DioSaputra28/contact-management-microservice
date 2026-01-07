package dto

type CreateContactRequest struct {
	UserId    string `json:"user_id" validate:"required"`
	FirstName string `json:"first_name" validate:"required,min=1,max=100"`
	LastName  string `json:"last_name" validate:"omitempty,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"omitempty,max=50"`
}

type UpdateContactRequest struct {
	FirstName string `json:"first_name" validate:"omitempty,min=1,max=100"`
	LastName  string `json:"last_name" validate:"omitempty,max=100"`
	Email     string `json:"email" validate:"omitempty,email"`
	Phone     string `json:"phone" validate:"omitempty,max=50"`
}
