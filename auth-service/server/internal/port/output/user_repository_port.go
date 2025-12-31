package output

import "github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/application/domain"

type UserRepositoryPort interface {
	FindByEmail(email string) (*domain.User, error)
	UpdateToken(userID int64, token string) error
	CreateUser(name, email, password string) (*domain.User, error)
}
