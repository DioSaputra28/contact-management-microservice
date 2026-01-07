package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Common error responses
var (
	ErrInternalServer = status.Error(codes.Internal, "internal server error")
	ErrNotFound       = status.Error(codes.NotFound, "resource not found")
)

// NewInvalidArgumentError creates a new InvalidArgument error
func NewInvalidArgumentError(message string) error {
	return status.Error(codes.InvalidArgument, message)
}

// NewNotFoundError creates a new NotFound error
func NewNotFoundError(message string) error {
	return status.Error(codes.NotFound, message)
}

// NewInternalError creates a new Internal error
func NewInternalError(message string) error {
	return status.Error(codes.Internal, message)
}
