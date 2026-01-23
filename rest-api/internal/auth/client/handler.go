package client

import (
	"github.com/DioSaputra28/contact-management-proto/protogen/go/auth"
	"google.golang.org/grpc"
)

type AuthHandlerClient struct {
	authInterface AuthInterface
}

func NewAuthHandlerClient(conn *grpc.ClientConn) *AuthHandlerClient {
	client := auth.NewAuthServiceClient(conn)
	
	return &AuthHandlerClient{
		authInterface: client,
	}
}