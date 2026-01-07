package output

import "github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/domain"

type ContactRepositoryPort interface {
	CreateContact(userId string, firstName, lastName, email, phone string) (*domain.Contact, error)
	UpdateContact(userId string, contactId int64, firstName, lastName, email, phone string) (*domain.Contact, error)
	DeleteContact(userId string, contactId int64) (*domain.Contact, error)
	GetContactById(userId string, contactId int64) (*domain.Contact, error)
	GetContacts(userId string, page, limit int, search string) ([]*domain.Contact, *domain.ContactPagination, error)
}
