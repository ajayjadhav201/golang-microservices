package auth

import (
	"api"
	"net/http"

	"github.com/ajayjadhav201/common"
)

type AuthClient struct {
	Client api.AuthServiceClient
}

func NewAuthClient(service api.AuthServiceClient) *AuthClient {
	return &AuthClient{service}
}

func (a *AuthClient) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v2/signup", a.SignupHandler)
	mux.HandleFunc("POST /api/v2/login", a.LoginHandler)
	mux.HandleFunc("POST /api/v2/changepassword", a.ChangePassword)
	mux.HandleFunc("POST /api/v2/updateuser", a.UpdateUserHandler)
	mux.HandleFunc("POST /api/v2/deleteuser", a.DeleteUser)
}

//
//
//
//
//

// Signup
func (a *AuthClient) SignupHandler(w http.ResponseWriter, r *http.Request) {
	req := &api.SignupRequest{}
	if err := common.ReadJSON(r, req); err != nil {
		common.Println("ajaj error while parsing json", err.Error())
		common.WriteRequestBodyError(w, err)
		return
	}
	//common.Println("ajaj signup request is ", req)

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
	var req *api.LoginRequest
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

func (a *AuthClient) ChangePassword(w http.ResponseWriter, r *http.Request) {
	//
	var req *api.ChangePasswordRequest
	if err := common.ReadJSON(r, req); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.ChangePassword(r.Context(), req)
	if err != nil {
		common.Println("ajaj chnage password errror is ", err)
		common.WriteGrpcError(w, err)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}

func (a *AuthClient) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	//
	var req *api.UpdateUserRequest
	if err := common.ReadJSON(r, req); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.UpdateUser(r.Context(), req)
	if err != nil {
		common.Println("ajaj updateuser errror is ", err)
		common.WriteGrpcError(w, err)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}

func (a *AuthClient) DeleteUser(w http.ResponseWriter, r *http.Request) {
	//
	var req *api.DeleteUserRequest
	if err := common.ReadJSON(r, req); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.DeleteUser(r.Context(), req)
	if err != nil {
		common.Println("ajaj deleteuser errror is ", err)
		common.WriteGrpcError(w, err)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}
