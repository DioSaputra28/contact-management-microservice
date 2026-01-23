package routes

import (
	"github.com/DioSaputra28/contact-management-microservice/rest-api/internal/auth/handler"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, authHandler *handler.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}
}
