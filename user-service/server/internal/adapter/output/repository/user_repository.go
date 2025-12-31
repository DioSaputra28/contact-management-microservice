package repository

import (
	"database/sql"

	"github.com/DioSaputra28/contact-management-microservice/user-service/server/internal/application/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(name, email, password string) (*domain.User, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := ur.db.Exec(query, name, email, password)
	if err != nil {
		return nil, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		UserID:   userID,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(id string, name, email, password string) (*domain.User, error) {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE user_id = ?"
	_, err := ur.db.Exec(query, name, email, password, id)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	return user, nil
}

func (ur *UserRepository) DeleteUser(id string) (*domain.User, error) {
	var user domain.User
	query := "SELECT user_id, name, email, password, created_at FROM users WHERE user_id = ?"
	err := ur.db.QueryRow(query, id).Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	deleteQuery := "DELETE FROM users WHERE user_id = ?"
	_, err = ur.db.Exec(deleteQuery, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserById(id int64) (*domain.User, error) {
	user := &domain.User{}
	query := "SELECT user_id, name, email, password, created_at FROM users WHERE user_id = ?"
	err := ur.db.QueryRow(query, id).Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUsers() ([]*domain.User, error) {
	query := "SELECT user_id, name, email, password, created_at FROM users"
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
