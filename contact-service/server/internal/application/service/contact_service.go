package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/domain"
	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/dto"
	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/validator"
	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/port/output"
)

type ContactService struct {
	contactRepo output.ContactRepositoryPort
}

func NewContactService(contactRepo output.ContactRepositoryPort) *ContactService {
	return &ContactService{contactRepo: contactRepo}
}

func (cs *ContactService) CreateContact(userId string, firstName, lastName, email, phone string) (*domain.Contact, error) {
	createReq := &dto.CreateContactRequest{
		UserId:    userId,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := validator.ValidateStruct(createReq); err != nil {
		return nil, err
	}

	return cs.contactRepo.CreateContact(userId, firstName, lastName, email, phone)
}

func (cs *ContactService) UpdateContact(userId string, contactId int64, firstName, lastName, email, phone string) (*domain.Contact, error) {
	updateReq := &dto.UpdateContactRequest{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := validator.ValidateStruct(updateReq); err != nil {
		return nil, err
	}

	return cs.contactRepo.UpdateContact(userId, contactId, firstName, lastName, email, phone)
}

func (cs *ContactService) DeleteContact(userId string, contactId int64) (*domain.Contact, error) {
	return cs.contactRepo.DeleteContact(userId, contactId)
}

func (cs *ContactService) GetContactById(userId string, contactId int64) (*domain.Contact, error) {
	contact, err := cs.contactRepo.GetContactById(userId, contactId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("contact not found")
		}
		return nil, err
	}

	return contact, nil
}

func (cs *ContactService) GetContacts(userId string, page, limit int, search string) ([]*domain.Contact, *domain.ContactPagination, error) {
	contacts, pagination, err := cs.contactRepo.GetContacts(userId, page, limit, search)
	fmt.Println(contacts)
	fmt.Println(err, "Di service")
	if err != nil {
		return nil, nil, err
	}

	return contacts, pagination, nil
}
