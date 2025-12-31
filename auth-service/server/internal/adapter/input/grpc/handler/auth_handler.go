package handler

import (
	"context"

	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/port/input"
	"github.com/DioSaputra28/contact-management-proto/protogen/go/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	authServicePort input.AuthServicePort
}

func NewAuthHandler(authServicePort input.AuthServicePort) *AuthHandler {
	return &AuthHandler{
		authServicePort: authServicePort,
	}
}

func (h *AuthHandler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, token, err := h.authServicePort.Login(req.Email, req.Password)
	if err != nil {
		switch err.Error() {
		case "invalid email or password":
			return nil, ErrInvalidCredentials
		case "validation failed: Email is required":
			return nil, NewInvalidArgumentError("email is required")
		case "validation failed: Password is required":
			return nil, NewInvalidArgumentError("password is required")
		default:
			if len(err.Error()) > 18 && err.Error()[:18] == "validation failed:" {
				return nil, NewInvalidArgumentError(err.Error())
			}
			return nil, ErrInternalServer
		}
	}

	return &auth.LoginResponse{
		Token: token,
		User: &auth.User{
			UserId:    int32(user.UserID),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}

func (h *AuthHandler) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	user, token, err := h.authServicePort.Register(req.Name, req.Email, req.Password)
	if err != nil {
		switch err.Error() {
		case "email already exists":
			return nil, ErrEmailAlreadyExists
		default:
			if len(err.Error()) > 18 && err.Error()[:18] == "validation failed:" {
				return nil, NewInvalidArgumentError(err.Error())
			}
			return nil, ErrInternalServer
		}
	}

	return &auth.RegisterResponse{
		Token: token,
		User: &auth.User{
			UserId:    int32(user.UserID),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}
