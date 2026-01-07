package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/config"
	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/adapter/input/grpc/handler"
	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/adapter/output/repository"
	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/service"
	"github.com/DioSaputra28/contact-management-proto/protogen/go/contact"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	// Database connection
	db, err := config.DbConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	// Initialize repositories
	contactRepo := repository.NewContactRepository(db)
	addressRepo := repository.NewAddressRepository(db)

	// Initialize services
	contactService := service.NewContactService(contactRepo)
	addressService := service.NewAddressService(addressRepo)

	// Initialize handler
	contactHandler := handler.NewContactHandler(contactService, addressService)

	// Setup gRPC server
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50053"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	contact.RegisterContactServiceServer(grpcServer, contactHandler)

	log.Printf("gRPC server starting on port %s...", port)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}
