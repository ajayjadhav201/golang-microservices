package routes

import (
	"context"
	"golang-microservices/api"
	"golang-microservices/common"

	"time"
	"user-service/database"
	"user-service/model"

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
	common.Println("ajaj signup request received email: ", req.Email, " and mobile number: ", req.MobileNumber)
	user := model.NewUser()

	err := common.Copy(req, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to read request body")
	}
	user.ID = common.UniqueID()
	tokn, err := common.CreateToken(common.Int64toa(user.ID))
	if err != nil {
		return nil, err
	}
	user.Token = tokn
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	//
	// calling create user database function
	id, err := s.db.CreateUser(user)
	//
	// handle errors
	if err != nil {
		return nil, err
	}

	// returning success result
	return &api.SignupResponse{
		Message: "user created successfully",
		Token:   id,
	}, nil
}

func (s *userService) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	return nil, nil
}
