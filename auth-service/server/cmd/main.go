package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/config"
	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/adapter/input/grpc/handler"
	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/adapter/output/repository"
	"github.com/DioSaputra28/contact-management-microservice/auth-service/server/internal/application/service"
	"github.com/DioSaputra28/contact-management-proto/protogen/go/auth"
	"google.golang.org/grpc"
)

func main() {
	db, err := config.DbConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authHandler)

	log.Printf("gRPC server starting on port %s...", port)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}
