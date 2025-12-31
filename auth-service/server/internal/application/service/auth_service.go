package service

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/application/domain"
	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/application/dto"
	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/application/validator"
	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/port/output"
)

type AuthService struct {
	userRepo output.UserRepositoryPort
}

func NewAuthService(userRepo output.UserRepositoryPort) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(email, password string) (*domain.User, string, error) {
	loginReq := &dto.LoginRequest{
		Email:    email,
		Password: password,
	}

	if err := validator.ValidateStruct(loginReq); err != nil {
		return nil, "", err
	}

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", errors.New("invalid email or password")
		}
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	token := uuid.New().String()
	err = s.userRepo.UpdateToken(user.UserID, token)
	if err != nil {
		return nil, "", err
	}

	user.Token = token

	return user, token, nil
}

func (s *AuthService) Register(name, email, password string) (*domain.User, string, error) {
	registerReq := &dto.RegisterRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := validator.ValidateStruct(registerReq); err != nil {
		return nil, "", err
	}

	user, err := s.userRepo.FindByEmail(email)
	if err != nil && err != sql.ErrNoRows {
		return nil, "", err
	}

	if user != nil {
		return nil, "", errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user, err = s.userRepo.CreateUser(name, email, string(hashedPassword))
	if err != nil {
		return nil, "", err
	}

	token := uuid.New().String()
	err = s.userRepo.UpdateToken(user.UserID, token)
	if err != nil {
		return nil, "", err
	}

	user.Token = token

	return user, token, nil
}
