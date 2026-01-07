package input

import "github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/domain"

type AddressServicePort interface {
	CreateAddress(contactId int64, street, city, state, zipCode, country string) (*domain.Address, error)
	UpdateAddress(contactId, addressId int64, street, city, state, zipCode, country string) (*domain.Address, error)
	DeleteAddress(contactId, addressId int64) (*domain.Address, error)
	GetAddressById(contactId, addressId int64) (*domain.Address, error)
	GetAddresses(contactId int64) ([]*domain.Address, error)
}
