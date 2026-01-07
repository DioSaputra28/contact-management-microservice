package service

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/DioSaputra28/contact-management-microservice/user-service/server/internal/application/domain"
	"github.com/DioSaputra28/contact-management-microservice/user-service/server/internal/application/dto"
	"github.com/DioSaputra28/contact-management-microservice/user-service/server/internal/application/validator"
	"github.com/DioSaputra28/contact-management-microservice/user-service/server/internal/port/output"
)

type UserService struct {
	userRepo output.UserRepositoryPort
}

func NewUserService(userRepo output.UserRepositoryPort) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) CreateUser(name, email, password string) (*domain.User, error) {
	createReq := &dto.CreateUserRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := validator.ValidateStruct(createReq); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return us.userRepo.CreateUser(name, email, string(hashedPassword))
}

func (us *UserService) UpdateUser(id string, name, email, password string) (*domain.User, error) {
	updateReq := &dto.UpdateUserRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := validator.ValidateStruct(updateReq); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return us.userRepo.UpdateUser(id, name, email, string(hashedPassword))
}

func (us *UserService) DeleteUser(id string) (*domain.User, error) {
	return us.userRepo.DeleteUser(id)
}

func (us *UserService) GetUserById(id int64) (*domain.User, error) {
	user, err := us.userRepo.GetUserById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetUsers(page, limit int, search string) ([]*domain.User, *domain.UserPagination, error) {
	users, pagination, err := us.userRepo.GetUsers(page, limit, search)
	if err != nil {
		return nil, nil, err
	}

	return users, pagination, nil
}
