package main

import (
	pb "github.com/ajayjadhav201/common/api"
) 

type UserService interface {
	Signup()
	Login()
}

type userService struct {
	pb.UnimplementedAuthServiceServer
}

func NewUserService() *userService {
	return &userService{}
}
