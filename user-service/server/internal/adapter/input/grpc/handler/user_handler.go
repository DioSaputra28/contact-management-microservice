package handler

import (
	"context"

	"github.com/DioSaputra28/contact-management-microservice/user-service/server/internal/port/input"
	"github.com/DioSaputra28/contact-management-proto/protogen/go/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	UserService input.UserServicePort
	user.UnimplementedUserServiceServer
}

func NewUserHandler(userS input.UserServicePort) *UserHandler {
	return &UserHandler{UserService: userS}
}

func (uh *UserHandler) GetUserById(ctx context.Context, req *user.GetUserByIdRequest) (*user.GetUserByIdResponse, error) {
	foundUser, err := uh.UserService.GetUserById(int64(req.UserId))
	if err != nil {
		return nil, err
	}

	return &user.GetUserByIdResponse{
		User: &user.User{
			UserId:    int32(foundUser.UserID),
			Name:      foundUser.Name,
			Email:     foundUser.Email,
			CreatedAt: timestamppb.New(*foundUser.CreatedAt),
		},
	}, nil
}

func (uh *UserHandler) GetUsers(ctx context.Context, req *user.GetUsersRequest) (*user.GetUsersResponse, error) {
	users, err := uh.UserService.GetUsers()
	if err != nil {
		return nil, err
	}

	var userResponses []*user.User
	for _, listUser := range users {
		userResponses = append(userResponses, &user.User{
			UserId:    int32(listUser.UserID),
			Name:      listUser.Name,
			Email:     listUser.Email,
			CreatedAt: timestamppb.New(*listUser.CreatedAt),
		})
	}

	return &user.GetUsersResponse{
		Users: userResponses,
	}, nil
}

func (uh *UserHandler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	createdUser, err := uh.UserService.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{
		User: &user.User{
			UserId:    int32(createdUser.UserID),
			Name:      createdUser.Name,
			Email:     createdUser.Email,
			CreatedAt: timestamppb.New(*createdUser.CreatedAt),
		},
	}, nil
}

func (uh *UserHandler) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	userID := string(req.UserId)

	updatedUser, err := uh.UserService.UpdateUser(userID, *req.Name, *req.Email, *req.Password)
	if err != nil {
		return nil, err
	}

	return &user.UpdateUserResponse{
		User: &user.User{
			UserId: req.UserId,
			Name:   updatedUser.Name,
			Email:  updatedUser.Email,
		},
	}, nil
}

func (uh *UserHandler) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	userID := string(req.UserId)

	deletedUser, err := uh.UserService.DeleteUser(userID)
	if err != nil {
		return nil, err
	}

	return &user.DeleteUserResponse{
		User: &user.User{
			UserId:    int32(deletedUser.UserID),
			Name:      deletedUser.Name,
			Email:     deletedUser.Email,
			CreatedAt: timestamppb.New(*deletedUser.CreatedAt),
		},
	}, nil
}
