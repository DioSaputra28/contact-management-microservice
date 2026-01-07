package dto

type CreateAddressRequest struct {
	ContactId int64  `json:"contact_id" validate:"required"`
	Street    string `json:"street" validate:"omitempty,max=255"`
	City      string `json:"city" validate:"omitempty,max=255"`
	State     string `json:"state" validate:"omitempty,max=255"`
	ZipCode   string `json:"zip_code" validate:"omitempty,max=20"`
	Country   string `json:"country" validate:"required,max=255"`
}

type UpdateAddressRequest struct {
	Street  string `json:"street" validate:"omitempty,max=255"`
	City    string `json:"city" validate:"omitempty,max=255"`
	State   string `json:"state" validate:"omitempty,max=255"`
	ZipCode string `json:"zip_code" validate:"omitempty,max=20"`
	Country string `json:"country" validate:"omitempty,max=255"`
}
