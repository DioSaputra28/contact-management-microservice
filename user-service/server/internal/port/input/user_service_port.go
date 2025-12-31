package input

import "github.com/DioSaputra28/contact-management-microservice/user-service/server/internal/application/domain"

type UserServicePort interface {
	GetUserById(id int64) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	CreateUser(name, email, password string) (*domain.User, error)
	UpdateUser(id string, name, email, password string) (*domain.User, error)
	DeleteUser(id string) (*domain.User, error)
}
