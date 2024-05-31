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
	Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)
	UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	ChangePassword(context.Context, *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error)
	DeleteUser(context.Context, *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
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
	common.Println("ajaj signup request received email: ", req.User.Email, " and mobile number: ", req.User.MobileNumber)
	user := model.NewUser()

	err := common.Copy(req.User, user)
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
	_, err = s.db.CreateUser(user)
	//
	// handle errors
	if err != nil {
		common.Println("ajaj printing error: ", err.Error())
		return nil, err
	}

	// returning success result
	return &pb.SignupResponse{
		Message: "user created successfully",
		Token:   user.Token,
		User:    req.User,
	}, nil
}

func (s *userService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// check if request body is not nil or empty
	if req == nil || req.EmailorMobile == "" || req.Password == "" {
		common.Println("ajaj login request invalid error: ", req.EmailorMobile, " pass: ", req.Password)
		return nil, common.Error("Login Rquest Invalid")
	}
	// if !common.IsEmail(req.EmailorMobile) && !common.IsMobileNumber(req.EmailorMobile) {
	// 	return nil, common.Error("Please enter a valid email or mobile")
	// }
	common.Println("ajaj login request with email:", req.EmailorMobile, " and password: ", req.Password)
	// find user in database
	user, err := s.db.GetUserByEmailorMobile(req.EmailorMobile)
	if err != nil {
		return nil, err
	}
	if user.Password != req.Password {
		common.Println("ajaj pass: ", user)
		common.Println("ajaj validpassword: ", user.Password, " and invalid pass: ", req.Password)
		return nil, common.Error("Please enter a valid password.")
	}

	return &pb.LoginResponse{
		Message: "user logged in successfully",
		Token:   user.Token,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	common.Println("ajaj update user request received")
	user := model.NewUser()

	err := common.Copy(req.User, user)
	if err != nil {
		return nil, common.Error("Unable to read request body")
	}
	user.ID = int64(common.Atoi(req.UserID))
	user.UpdatedAt = time.Now()
	updatedUser, err := s.db.UpdateUser(req.UserID, user)
	if err != nil {
		return nil, err
	}

	common.Println("updated user is: ", updatedUser)

	return &pb.UpdateUserResponse{
		Message: "User updated successfully.",
		Token:   "na",
		User:    req.User,
	}, nil
}

//
//

func (s *userService) ChangePassword(context.Context, *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}

//
//

func (s *userService) DeleteUser(context.Context, *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
