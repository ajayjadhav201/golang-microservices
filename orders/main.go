package main

import (
	"context"
	"net"

	"github.com/ajayjadhav201/common"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	common.Fatal(err)
	defer l.Close()

	//
	store := NewStore()
	svc := NewService(store) //this is for connection with database
	// NewGrpcHandler(grpcServer)

	svc.CreateOrder(context.Background())
	common.Println("Order service is started on: ", grpcAddr)
	common.Fatal(grpcServer.Serve(l))
}
