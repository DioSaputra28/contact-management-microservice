package handler

import (
	"net/http"

	"github.com/DioSaputra28/contact-management-microservice/rest-api/internal/auth/client"
	"github.com/DioSaputra28/contact-management-microservice/rest-api/internal/auth/dto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	authClient *client.AuthClient
}

func NewAuthHandler(authClient *client.AuthClient) *AuthHandler {
	return &AuthHandler{
		authClient: authClient,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	resp, err := h.authClient.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   "internal_error",
				Message: "An unexpected error occurred",
			})
			return
		}

		switch st.Code() {
		case codes.NotFound:
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error:   "not_found",
				Message: st.Message(),
			})
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "invalid_argument",
				Message: st.Message(),
			})
		case codes.Unauthenticated:
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Error:   "unauthorized",
				Message: st.Message(),
			})
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   "internal_error",
				Message: "An unexpected error occurred",
			})
		}
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		Token: resp.Token,
		User: dto.UserData{
			UserID: resp.User.UserId,
			Name:   resp.User.Name,
			Email:  resp.User.Email,
		},
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	resp, err := h.authClient.Register(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   "internal_error",
				Message: "An unexpected error occurred",
			})
			return
		}

		switch st.Code() {
		case codes.AlreadyExists:
			c.JSON(http.StatusConflict, dto.ErrorResponse{
				Error:   "already_exists",
				Message: st.Message(),
			})
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "invalid_argument",
				Message: st.Message(),
			})
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   "internal_error",
				Message: "An unexpected error occurred",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, dto.AuthResponse{
		Token: resp.Token,
		User: dto.UserData{
			UserID: resp.User.UserId,
			Name:   resp.User.Name,
			Email:  resp.User.Email,
		},
	})
}
