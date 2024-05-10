package auth

import (
	"net/http"

	pb "github.com/ajayjadhav201/common/api"
)

type Authentication struct {
	Client pb.AuthServiceClient
}

func NewAuthClient(service pb.AuthServiceClient) *Authentication {
	return &Authentication{
		service}
}

func (a *Authentication) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v2/signup", a.SignupHandler)
	mux.HandleFunc("POST /api/v2/login", a.LoginHandler)
}
