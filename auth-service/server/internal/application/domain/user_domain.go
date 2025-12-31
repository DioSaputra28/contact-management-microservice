package domain

import "time"

type User struct {
	UserID    int64     `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	Token     string    `db:"token" json:"token"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
