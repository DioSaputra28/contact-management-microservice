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
		return nil, err
	}

	return &auth.LoginResponse{
		Token: token,
		User: &auth.User{
			UserId:    int32(user.UserID),
			Name:      "",
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}

func (h *AuthHandler) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	user, token, err := h.authServicePort.Register(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		Token: token,
		User: &auth.User{
			UserId:    int32(user.UserID),
			Name:      "",
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}
