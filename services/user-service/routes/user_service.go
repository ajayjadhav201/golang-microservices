package routes

import (
	"context"
	"golang-microservices/api/pb"
	"golang-microservices/common"

	"time"
	"user-service/database"
	"user-service/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Signup(context.Context, *pb.SignupRequest) (*pb.SignupResponse, error)
	// Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)
}

type userService struct {
	pb.UnimplementedAuthServiceServer
	db database.UserStore
}

func NewUserService(store database.UserStore) *userService {
	return &userService{
		db: store,
	}
}

func (s *userService) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
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
	return &pb.SignupResponse{
		Message: "user created successfully",
		Token:   id,
	}, nil
}

func (s *userService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}
