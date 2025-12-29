package domain

import "time"

type User struct {
	UserID    int64     `db:"user_id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
