package auth

import (
	"net/http"

	"github.com/ajayjadhav201/common"
	pb "github.com/ajayjadhav201/common/api"
)

// Signup
func (a *Authentication) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req *pb.SignupRequest
	if err := common.ReadJSON(r, req); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.Signup(r.Context(), req)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, common.InternalServerErr)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}

// Login
func (a *Authentication) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req *pb.LoginRequest
	if err := common.ReadJSON(r, req); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.Login(r.Context(), req)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, common.InternalServerErr)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}
