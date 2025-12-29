package input

import "github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/application/domain"

type AuthServicePort interface {
	Login(email, password string) (*domain.User, string, error)
	Register(name, email, password string) (string, error)
}
