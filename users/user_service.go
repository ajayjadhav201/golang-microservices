package main

import (
	"api"
	"context"
	"model"
	"time"

	"github.com/ajayjadhav201/common"
	"github.com/ajayjadhav201/users/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Signup(context.Context, *api.SignupRequest) (*api.SignupResponse, error)
	// Login(context.Context, *api.LoginRequest) (*api.LoginResponse, error)
}

type userService struct {
	api.UnimplementedAuthServiceServer
	db database.UserStore
}

func NewUserService(store database.UserStore) *userService {
	return &userService{
		db: store,
	}
}

func (s *userService) Signup(ctx context.Context, req *api.SignupRequest) (*api.SignupResponse, error) {
	common.Println("ajaj signup request received ", req)
	user := model.NewUser()

	err := common.Copy(req, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to read request body")
	}
	tokn, err := common.CreateToken(req.Email)
	if err != nil {
		return nil, err
	}
	user.ID = common.UniqueID()
	user.Token = tokn
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	common.Println("ajaj user before updating to database: ", user)
	id, err := s.db.CreateUser(user)
	if err != nil {
		common.Println("Error while creating user: ", err)
		return nil, err
	}
	common.Println("ajaj user object is ", user)

	return &api.SignupResponse{
		Message: "user create successfully",
		Token:   id,
	}, nil
}

func (s *userService) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	return nil, nil
}
