package repository

import (
	"database/sql"
	"time"

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

	now := time.Now()
	user := &domain.User{
		UserID:    userID,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: &now,
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(id string, name, email, password string) (*domain.User, error) {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE user_id = ?"
	_, err := ur.db.Exec(query, name, email, password, id)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	user := &domain.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: &now,
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

func (ur *UserRepository) GetUsers(page, limit int, search string) ([]*domain.User, *domain.UserPagination, error) {
	var users []*domain.User
	var totalData int64

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10 
	}
	if limit > 100 {
		limit = 100 
	}

	query := `
		SELECT 
			user_id, 
			name, 
			email, 
			password, 
			created_at,
			COUNT(*) OVER() as total_count
		FROM users
	`

	var args []any

	if search != "" {
		query += " WHERE name LIKE ? OR email LIKE ?"
		searchPattern := "%" + search + "%"
		args = append(args, searchPattern, searchPattern)
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, (page-1)*limit)

	rows, err := ur.db.Query(query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(
			&user.UserID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&totalData,
		)
		if err != nil {
			return nil, nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 && totalData == 0 {
		countQuery := "SELECT COUNT(*) FROM users"
		var countArgs []any

		if search != "" {
			countQuery += " WHERE name LIKE ? OR email LIKE ?"
			searchPattern := "%" + search + "%"
			countArgs = append(countArgs, searchPattern, searchPattern)
		}

		err := ur.db.QueryRow(countQuery, countArgs...).Scan(&totalData)
		if err != nil {
			return nil, nil, err
		}
	}

	totalPage := int64(0)
	if totalData > 0 {
		totalPage = totalData / int64(limit)
		if totalData%int64(limit) > 0 {
			totalPage++
		}
	}

	return users, &domain.UserPagination{
		TotalData:   totalData,
		CurrentPage: int64(page),
		PageSize:    int64(limit),
		TotalPage:   totalPage,
	}, nil
}
