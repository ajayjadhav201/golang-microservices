package main

import (
	"api"
	"net/http"

	"github.com/ajayjadhav201/common"
	"github.com/ajayjadhav201/gateway/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGrpcClient(mux *http.ServeMux) {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	//
	// orderClient, err := grpc.Dial(orderServiceAddr, creds)
	// common.Fatal(err)
	// defer orderClient.Close()
	//S
	authClient, err := grpc.NewClient("localhost:2001", creds)
	common.Fatal(err)

	//
	// handle order client
	// NewHandler(pb.NewOrderServiceClient(orderClient)).registerRoutes(mux)
	//
	// handle auth client
	auth.NewAuthClient(api.NewAuthServiceClient(authClient)).RegisterRoutes(mux)
}
