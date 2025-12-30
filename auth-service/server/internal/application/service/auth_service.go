package service

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/application/domain"
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

func (s *AuthService) Register(email, password string) (*domain.User, string, error) {
	// Check if email already exists
	user, err := s.userRepo.FindByEmail(email)
	if err != nil && err != sql.ErrNoRows {
		// Return error only if it's NOT "no rows" error
		return nil, "", err
	}

	if user != nil {
		// Email already exists
		return nil, "", errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user, err = s.userRepo.CreateUser(email, string(hashedPassword))
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
