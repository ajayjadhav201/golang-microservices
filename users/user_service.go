package main

import (
	"api"
	"context"
)

type UserService interface {
	Signup(context.Context, *api.SignupRequest) (*api.SignupResponse, error)
	Login(context.Context, *api.LoginRequest) (*api.LoginResponse, error)
}

type userService struct {
	api.UnimplementedAuthServiceServer
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) Signup(context.Context, *api.SignupRequest) (*api.SignupResponse, error) {
	return nil, nil
}

func (s *userService) Login(context.Context, *api.LoginRequest) (*api.LoginResponse, error) {
	return nil, nil
}
