package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Common error responses
var (
	ErrEmailAlreadyExists = status.Error(codes.AlreadyExists, "email already exists")
	ErrInternalServer     = status.Error(codes.Internal, "internal server error")
)

// NewInvalidArgumentError creates a new InvalidArgument error
func NewInvalidArgumentError(message string) error {
	return status.Error(codes.InvalidArgument, message)
}

// NewNotFoundError creates a new NotFound error
func NewNotFoundError(message string) error {
	return status.Error(codes.NotFound, message)
}

// NewUnauthenticatedError creates a new Unauthenticated error
func NewUnauthenticatedError(message string) error {
	return status.Error(codes.Unauthenticated, message)
}

// NewAlreadyExistsError creates a new AlreadyExists error
func NewAlreadyExistsError(message string) error {
	return status.Error(codes.AlreadyExists, message)
}

// NewInternalError creates a new Internal error
func NewInternalError(message string) error {
	return status.Error(codes.Internal, message)
}
