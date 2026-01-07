package service

import (
	"database/sql"
	"errors"

	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/domain"
	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/dto"
	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/validator"
	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/port/output"
)

type AddressService struct {
	addressRepo output.AddressRepositoryPort
}

func NewAddressService(addressRepo output.AddressRepositoryPort) *AddressService {
	return &AddressService{addressRepo: addressRepo}
}

func (as *AddressService) CreateAddress(contactId int64, street, city, state, zipCode, country string) (*domain.Address, error) {
	createReq := &dto.CreateAddressRequest{
		ContactId: contactId,
		Street:    street,
		City:      city,
		State:     state,
		ZipCode:   zipCode,
		Country:   country,
	}

	if err := validator.ValidateStruct(createReq); err != nil {
		return nil, err
	}

	return as.addressRepo.CreateAddress(contactId, street, city, state, zipCode, country)
}

func (as *AddressService) UpdateAddress(contactId, addressId int64, street, city, state, zipCode, country string) (*domain.Address, error) {
	updateReq := &dto.UpdateAddressRequest{
		Street:  street,
		City:    city,
		State:   state,
		ZipCode: zipCode,
		Country: country,
	}

	if err := validator.ValidateStruct(updateReq); err != nil {
		return nil, err
	}

	return as.addressRepo.UpdateAddress(contactId, addressId, street, city, state, zipCode, country)
}

func (as *AddressService) DeleteAddress(contactId, addressId int64) (*domain.Address, error) {
	return as.addressRepo.DeleteAddress(contactId, addressId)
}

func (as *AddressService) GetAddressById(contactId, addressId int64) (*domain.Address, error) {
	address, err := as.addressRepo.GetAddressById(contactId, addressId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("address not found")
		}
		return nil, err
	}

	return address, nil
}

func (as *AddressService) GetAddresses(contactId int64) ([]*domain.Address, error) {
	addresses, err := as.addressRepo.GetAddresses(contactId)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}
