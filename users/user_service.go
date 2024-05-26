package main

import (
	"api"
	"context"

	"github.com/ajayjadhav201/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Signup(context.Context, *api.SignupRequest) (*api.SignupResponse, error)
	Login(context.Context, *api.LoginRequest) (*api.LoginResponse, error)
}

type userService struct {
	api.UnimplementedAuthServiceServer
}

func NewUserService() *userService {
	return &userService{}
}

func (s *userService) Signup(ctx context.Context, req *api.SignupRequest) (*api.SignupResponse, error) {
	common.Println("ajaj signup request received ", req)

	return nil, status.Errorf(codes.Internal, "Internal error")
}

func (s *userService) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	return nil, nil
}
