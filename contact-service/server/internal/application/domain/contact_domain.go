package domain

import "time"

type Contact struct {
	ContactId int        `db:"contact_id" json:"contact_id"`
	UserId    string     `db:"user_id" json:"user_id"`
	FirstName string     `db:"first_name" json:"first_name"`
	LastName  string     `db:"last_name" json:"last_name"`
	Email     string     `db:"email" json:"email"`
	Phone     string     `db:"phone" json:"phone"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}

type ContactPagination struct {
	TotalData   int64 `json:"total_data"`
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
	TotalPage   int64 `json:"total_page"`
}
