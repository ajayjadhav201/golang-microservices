package main

import (
	"net/http"

	"github.com/ajayjadhav201/common"
	pb "github.com/ajayjadhav201/common/api"
	"github.com/ajayjadhav201/gateway/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GrpcClient(mux *http.ServeMux) {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	//
	orderClient, err := grpc.Dial(orderServiceAddr, creds)
	common.Fatal(err)
	// defer orderClient.Close()
	//
	authClient, err := grpc.Dial(authServiceAddr, creds)
	common.Fatal(err)
	defer authClient.Close()
	//
	// handle product client
	NewHandler(pb.NewOrderServiceClient(orderClient)).registerRoutes(mux)
	//
	// handle auth client
	auth.NewAuthClient(pb.NewAuthServiceClient(authClient)).Register(mux)
}
