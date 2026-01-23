package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/port/input"
	"github.com/DioSaputra28/contact-management-proto/protogen/go/contact"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ContactHandler struct {
	contactService input.ContactServicePort
	addressService input.AddressServicePort
	contact.UnimplementedContactServiceServer
}

func NewContactHandler(contactService input.ContactServicePort, addressService input.AddressServicePort) *ContactHandler {
	return &ContactHandler{
		contactService: contactService,
		addressService: addressService,
	}
}

// Contact operations

func (h *ContactHandler) GetContactById(ctx context.Context, req *contact.GetContactByIdRequest) (*contact.GetContactByIdResponse, error) {
	userId := strconv.Itoa(int(req.GetUserId()))
	contactId := int64(req.GetContactId())

	foundContact, err := h.contactService.GetContactById(userId, contactId)
	if err != nil {
		if err.Error() == "contact not found" {
			return nil, NewNotFoundError("contact not found")
		}
		if len(err.Error()) > 18 && err.Error()[:18] == "validation failed:" {
			return nil, NewInvalidArgumentError(err.Error())
		}
		return nil, ErrInternalServer
	}

	lastName := foundContact.LastName
	phone := foundContact.Phone

	return &contact.GetContactByIdResponse{
		Contact: &contact.Contact{
			ContactId: int32(foundContact.ContactId),
			UserId:    int32(req.GetUserId()),
			FirstName: foundContact.FirstName,
			LastName:  &lastName,
			Email:     foundContact.Email,
			Phone:     &phone,
			CreatedAt: timestamppb.New(*foundContact.CreatedAt),
		},
	}, nil
}

func (h *ContactHandler) GetContacts(ctx context.Context, req *contact.GetContactsRequest) (*contact.GetContactsResponse, error) {
	userId := strconv.Itoa(int(req.GetUserId()))
	page := int(req.GetPage())
	limit := int(req.GetLimit())
	search := ""
	if req.Search != nil {
		search = *req.Search
	}

	contacts, pagination, err := h.contactService.GetContacts(userId, page, limit, search)
	if err != nil {
		return nil, ErrInternalServer
	}

	fmt.Println("Sampai line 73")

	var contactResponses []*contact.Contact
	for _, c := range contacts {
		lastName := c.LastName
		phone := c.Phone

		contactResponses = append(contactResponses, &contact.Contact{
			ContactId: int32(c.ContactId),
			UserId:    req.GetUserId(),
			FirstName: c.FirstName,
			LastName:  &lastName,
			Email:     c.Email,
			Phone:     &phone,
			CreatedAt: timestamppb.New(*c.CreatedAt),
		})
	}

	return &contact.GetContactsResponse{
		Contacts:    contactResponses,
		TotalData:   int32(pagination.TotalData),
		CurrentPage: int32(pagination.CurrentPage),
		PageSize:    int32(pagination.PageSize),
		TotalPage:   int32(pagination.TotalPage),
	}, nil
}

func (h *ContactHandler) CreateContact(ctx context.Context, req *contact.CreateContactRequest) (*contact.CreateContactResponse, error) {
	userId := strconv.Itoa(int(req.GetUserId()))
	firstName := req.GetFirstName()
	lastName := ""
	if req.LastName != nil {
		lastName = *req.LastName
	}
	email := req.GetEmail()
	phone := ""
	if req.Phone != nil {
		phone = *req.Phone
	}

	createdContact, err := h.contactService.CreateContact(userId, firstName, lastName, email, phone)
	if err != nil {
		if len(err.Error()) > 18 && err.Error()[:18] == "validation failed:" {
			return nil, NewInvalidArgumentError(err.Error())
		}
		return nil, ErrInternalServer
	}

	lastNamePtr := createdContact.LastName
	phonePtr := createdContact.Phone

	return &contact.CreateContactResponse{
		Contact: &contact.Contact{
			ContactId: int32(createdContact.ContactId),
			UserId:    req.GetUserId(),
			FirstName: createdContact.FirstName,
			LastName:  &lastNamePtr,
			Email:     createdContact.Email,
			Phone:     &phonePtr,
			CreatedAt: timestamppb.New(*createdContact.CreatedAt),
		},
	}, nil
}

func (h *ContactHandler) UpdateContact(ctx context.Context, req *contact.UpdateContactRequest) (*contact.UpdateContactResponse, error) {
	userId := strconv.Itoa(int(req.GetUserId()))
	contactId := int64(req.GetContactId())
	firstName := ""
	if req.FirstName != nil {
		firstName = *req.FirstName
	}
	lastName := ""
	if req.LastName != nil {
		lastName = *req.LastName
	}
	email := ""
	if req.Email != nil {
		email = *req.Email
	}
	phone := ""
	if req.Phone != nil {
		phone = *req.Phone
	}

	updatedContact, err := h.contactService.UpdateContact(userId, contactId, firstName, lastName, email, phone)
	if err != nil {
		if err.Error() == "contact not found" {
			return nil, NewNotFoundError("contact not found")
		}
		if len(err.Error()) > 18 && err.Error()[:18] == "validation failed:" {
			return nil, NewInvalidArgumentError(err.Error())
		}
		return nil, ErrInternalServer
	}

	lastNamePtr := updatedContact.LastName
	phonePtr := updatedContact.Phone

	return &contact.UpdateContactResponse{
		Contact: &contact.Contact{
			ContactId: int32(updatedContact.ContactId),
			UserId:    req.GetUserId(),
			FirstName: updatedContact.FirstName,
			LastName:  &lastNamePtr,
			Email:     updatedContact.Email,
			Phone:     &phonePtr,
			CreatedAt: timestamppb.New(*updatedContact.CreatedAt),
		},
	}, nil
}

func (h *ContactHandler) DeleteContact(ctx context.Context, req *contact.DeleteContactRequest) (*contact.DeleteContactResponse, error) {
	userId := strconv.Itoa(int(req.GetUserId()))
	contactId := int64(req.GetContactId())

	deletedContact, err := h.contactService.DeleteContact(userId, contactId)
	if err != nil {
		if err.Error() == "contact not found" {
			return nil, NewNotFoundError("contact not found")
		}
		return nil, ErrInternalServer
	}

	lastNamePtr := deletedContact.LastName
	phonePtr := deletedContact.Phone

	return &contact.DeleteContactResponse{
		Contact: &contact.Contact{
			ContactId: int32(deletedContact.ContactId),
			UserId:    req.GetUserId(),
			FirstName: deletedContact.FirstName,
			LastName:  &lastNamePtr,
			Email:     deletedContact.Email,
			Phone:     &phonePtr,
			CreatedAt: timestamppb.New(*deletedContact.CreatedAt),
		},
	}, nil
}
