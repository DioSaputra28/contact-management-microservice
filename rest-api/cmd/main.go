package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DioSaputra28/contact-management-microservice/rest-api/config"
	"github.com/DioSaputra28/contact-management-microservice/rest-api/internal/auth/client"
	"github.com/DioSaputra28/contact-management-microservice/rest-api/internal/auth/handler"
	"github.com/DioSaputra28/contact-management-microservice/rest-api/internal/auth/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	authConn := config.NewAuthServiceConnection()
	defer authConn.Close()

	authClient := client.NewAuthClient(authConn)

	authHandler := handler.NewAuthHandler(authClient)

	r := gin.Default()

	api := r.Group("/api/v1")

	routes.RegisterAuthRoutes(api, authHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("REST API Server running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
