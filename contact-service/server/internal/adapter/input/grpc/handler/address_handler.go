package handler

import (
	"context"

	"github.com/DioSaputra28/contact-management-proto/protogen/go/contact"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Address operations

func (h *ContactHandler) GetAddressById(ctx context.Context, req *contact.GetAddressByIdRequest) (*contact.GetAddressByIdResponse, error) {
	contactId := int64(req.GetContactId())
	addressId := int64(req.GetAddressId())

	foundAddress, err := h.addressService.GetAddressById(contactId, addressId)
	if err != nil {
		if err.Error() == "address not found" {
			return nil, NewNotFoundError("address not found")
		}
		return nil, ErrInternalServer
	}

	street := foundAddress.Street
	city := foundAddress.City
	state := foundAddress.State
	zipCode := foundAddress.ZipCode

	return &contact.GetAddressByIdResponse{
		Address: &contact.Address{
			AddressId: int32(foundAddress.AddressId),
			ContactId: int32(foundAddress.ContactId),
			Street:    &street,
			City:      &city,
			State:     &state,
			ZipCode:   &zipCode,
			Country:   foundAddress.Country,
			CreatedAt: timestamppb.New(*foundAddress.CreatedAt),
		},
	}, nil
}

func (h *ContactHandler) GetAddresses(ctx context.Context, req *contact.GetAddressesRequest) (*contact.GetAddressesResponse, error) {
	contactId := int64(req.GetContactId())

	addresses, err := h.addressService.GetAddresses(contactId)
	if err != nil {
		return nil, ErrInternalServer
	}

	var addressResponses []*contact.Address
	for _, a := range addresses {
		street := a.Street
		city := a.City
		state := a.State
		zipCode := a.ZipCode

		addressResponses = append(addressResponses, &contact.Address{
			AddressId: int32(a.AddressId),
			ContactId: int32(a.ContactId),
			Street:    &street,
			City:      &city,
			State:     &state,
			ZipCode:   &zipCode,
			Country:   a.Country,
			CreatedAt: timestamppb.New(*a.CreatedAt),
		})
	}

	return &contact.GetAddressesResponse{
		Addresses: addressResponses,
	}, nil
}

func (h *ContactHandler) CreateAddress(ctx context.Context, req *contact.CreateAddressRequest) (*contact.CreateAddressResponse, error) {
	contactId := int64(req.GetContactId())
	street := req.GetStreet()
	city := req.GetCity()
	state := req.GetState()
	zipCode := req.GetZipCode()
	country := req.GetCountry()

	createdAddress, err := h.addressService.CreateAddress(contactId, street, city, state, zipCode, country)
	if err != nil {
		if len(err.Error()) > 18 && err.Error()[:18] == "validation failed:" {
			return nil, NewInvalidArgumentError(err.Error())
		}
		return nil, ErrInternalServer
	}

	streetPtr := createdAddress.Street
	cityPtr := createdAddress.City
	statePtr := createdAddress.State
	zipCodePtr := createdAddress.ZipCode

	return &contact.CreateAddressResponse{
		Address: &contact.Address{
			AddressId: int32(createdAddress.AddressId),
			ContactId: int32(createdAddress.ContactId),
			Street:    &streetPtr,
			City:      &cityPtr,
			State:     &statePtr,
			ZipCode:   &zipCodePtr,
			Country:   createdAddress.Country,
			CreatedAt: timestamppb.New(*createdAddress.CreatedAt),
		},
	}, nil
}

func (h *ContactHandler) UpdateAddress(ctx context.Context, req *contact.UpdateAddressRequest) (*contact.UpdateAddressResponse, error) {
	contactId := int64(req.GetContactId())
	addressId := int64(req.GetAddressId())
	street := ""
	if req.Street != nil {
		street = *req.Street
	}
	city := ""
	if req.City != nil {
		city = *req.City
	}
	state := ""
	if req.State != nil {
		state = *req.State
	}
	zipCode := ""
	if req.ZipCode != nil {
		zipCode = *req.ZipCode
	}
	country := ""
	if req.Country != nil {
		country = *req.Country
	}

	updatedAddress, err := h.addressService.UpdateAddress(contactId, addressId, street, city, state, zipCode, country)
	if err != nil {
		if err.Error() == "address not found" {
			return nil, NewNotFoundError("address not found")
		}
		if len(err.Error()) > 18 && err.Error()[:18] == "validation failed:" {
			return nil, NewInvalidArgumentError(err.Error())
		}
		return nil, ErrInternalServer
	}

	streetPtr := updatedAddress.Street
	cityPtr := updatedAddress.City
	statePtr := updatedAddress.State
	zipCodePtr := updatedAddress.ZipCode

	return &contact.UpdateAddressResponse{
		Address: &contact.Address{
			AddressId: int32(updatedAddress.AddressId),
			ContactId: int32(updatedAddress.ContactId),
			Street:    &streetPtr,
			City:      &cityPtr,
			State:     &statePtr,
			ZipCode:   &zipCodePtr,
			Country:   updatedAddress.Country,
			CreatedAt: timestamppb.New(*updatedAddress.CreatedAt),
		},
	}, nil
}

func (h *ContactHandler) DeleteAddress(ctx context.Context, req *contact.DeleteAddressRequest) (*contact.DeleteAddressResponse, error) {
	contactId := int64(req.GetContactId())
	addressId := int64(req.GetAddressId())

	deletedAddress, err := h.addressService.DeleteAddress(contactId, addressId)
	if err != nil {
		if err.Error() == "address not found" {
			return nil, NewNotFoundError("address not found")
		}
		return nil, ErrInternalServer
	}

	streetPtr := deletedAddress.Street
	cityPtr := deletedAddress.City
	statePtr := deletedAddress.State
	zipCodePtr := deletedAddress.ZipCode

	return &contact.DeleteAddressResponse{
		Address: &contact.Address{
			AddressId: int32(deletedAddress.AddressId),
			ContactId: int32(deletedAddress.ContactId),
			Street:    &streetPtr,
			City:      &cityPtr,
			State:     &statePtr,
			ZipCode:   &zipCodePtr,
			Country:   deletedAddress.Country,
			CreatedAt: timestamppb.New(*deletedAddress.CreatedAt),
		},
	}, nil
}
