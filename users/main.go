package main

import (
	"net"

	"github.com/ajayjadhav201/common"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("USER_SERVICE_ADDRESS", "localhost:2001")
)

func main() {

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	common.Fatal(err)
	defer l.Close()

	// NewGrpcHandler(grpcServer)

	common.Println("Auth service is started on: ", grpcAddr)
	common.Fatal(grpcServer.Serve(l))
}
