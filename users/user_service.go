package main

import (
	"api"
)

type UserService interface {
	Signup()
	Login()
}

type userService struct {
	api.UnimplementedAuthServiceServer
}

func NewUserService() *userService {
	return &userService{}
}
