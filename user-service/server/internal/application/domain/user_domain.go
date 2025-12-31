package domain

import "time"

type User struct {
	UserID    int64      `json:"user_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
}
