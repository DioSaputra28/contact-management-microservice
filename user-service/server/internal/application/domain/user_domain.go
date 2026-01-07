package domain

import "time"

type User struct {
	UserID    int64      `json:"user_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
}

type UserPagination struct {
	TotalData int64 `json:"total_data"`
	CurrentPage int64 `json:"current_page"`
	PageSize int64 `json:"page_size"`
	TotalPage int64 `json:"total_page"`
}