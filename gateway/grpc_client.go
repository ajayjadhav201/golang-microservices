package main

import (
	"net/http"

	"github.com/ajayjadhav201/common"
	pb "github.com/ajayjadhav201/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGrpcClient(mux *http.ServeMux) {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	//
	// orderClient, err := grpc.Dial(orderServiceAddr, creds)
	// common.Fatal(err)
	// defer orderClient.Close()
	//
	authClient, err := grpc.Dial("localhost:2001", creds)
	common.Fatal(err)
	common.Println("ajaj errro while connecting to authservice ", err)

	//
	// handle order client
	// NewHandler(pb.NewOrderServiceClient(orderClient)).registerRoutes(mux)
	//
	// handle auth client
	NewAuthClient(pb.NewAuthServiceClient(authClient)).RegisterRoutes(mux)
}
