package config

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthServiceConnection() *grpc.ClientConn {
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		authServiceURL = "localhost:50051"
	}

	conn, err := grpc.NewClient(authServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to auth-service: %v", err)
	}

	fmt.Printf("Connected to auth-service at %s\n", authServiceURL)
	return conn
}
