package client

import (
	"context"

	"github.com/DioSaputra28/contact-management-proto/protogen/go/auth"
	"google.golang.org/grpc"
)

type AuthClient struct {
	client auth.AuthServiceClient
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	return &AuthClient{
		client: auth.NewAuthServiceClient(conn),
	}
}

func (ac *AuthClient) Login(ctx context.Context, email, password string) (*auth.LoginResponse, error) {
	req := &auth.LoginRequest{
		Email:    email,
		Password: password,
	}
	return ac.client.Login(ctx, req)
}

func (ac *AuthClient) Register(ctx context.Context, name, email, password string) (*auth.RegisterResponse, error) {
	req := &auth.RegisterRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return ac.client.Register(ctx, req)
}
