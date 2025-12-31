package repository

import (
	"database/sql"

	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/application/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}

	query := "SELECT user_id, name, email, password, token, created_at, updated_at FROM users WHERE email = ?"
	err := r.db.QueryRow(query, email).Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Token,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateToken(userID int64, token string) error {
	query := "UPDATE users SET token = ? WHERE user_id = ?"
	_, err := r.db.Exec(query, token, userID)
	return err
}

func (r *UserRepository) CreateUser(name, email, password string) (*domain.User, error) {
	user := &domain.User{}

	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, name, email, password)
	if err != nil {
		return nil, err
	}

	user.UserID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.Email = email
	user.Name = name

	return user, nil
}
