package client

import (
	"context"

	"github.com/DioSaputra28/contact-management-proto/protogen/go/auth"
	"google.golang.org/grpc"
)

type AuthInterface interface {
	Login(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.LoginResponse, error)
	Register(ctx context.Context, in *auth.RegisterRequest, opts ...grpc.CallOption) (*auth.RegisterResponse, error)
}