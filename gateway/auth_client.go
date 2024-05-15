package main

import (
	"net/http"

	"github.com/ajayjadhav201/common"
	pb "github.com/ajayjadhav201/common/api"
)

type AuthClient struct {
	Client pb.AuthServiceClient
}

func NewAuthClient(service pb.AuthServiceClient) *AuthClient {
	return &AuthClient{service}
}

func (a *AuthClient) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v2/signup", a.SignupHandler)
	mux.HandleFunc("POST /api/v2/login", a.LoginHandler)
	mux.HandleFunc("POST /api/v2/setpassword", a.SetPassword)
	mux.HandleFunc("POST /api/v2/updateuser", a.UpdateUserHandler)
}

//
//
//
//
//

// Signup
func (a *AuthClient) SignupHandler(w http.ResponseWriter, r *http.Request) {
	req := &pb.SignupRequest{}
	if err := common.ReadJSON(r, req); err != nil {
		common.Println("ajaj error while parsing json", err.Error())
		common.WriteError(w, http.StatusBadRequest, common.SPrintf("aj error %s", err.Error()))
		return
	}
	common.Println("ajaj signup request is ", req)

	res, err := a.Client.Signup(r.Context(), req)
	if err != nil {
		common.Println("ajaj signup errror is ", err)
		common.WriteGrpcError(w, err)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}

// Login
func (a *AuthClient) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req *pb.LoginRequest
	if err := common.ReadJSON(r, req); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.Login(r.Context(), req)
	if err != nil {
		common.Println("ajaj login errror is ", err)
		common.WriteGrpcError(w, err)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}

func (a *AuthClient) ForgotPassword() {
	//
}

func (a *AuthClient) SetPassword(w http.ResponseWriter, r *http.Request) {
	//
}

func (a *AuthClient) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	//
}
