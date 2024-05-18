package main

import (
	"net"

	"github.com/ajayjadhav201/common"
	"github.com/ajayjadhav201/users/database"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("USER_SERVICE_ADDRESS", "localhost:2001")
)

func main() {
	err := godotenv.Load(".env")
	common.Fatal(err)

	//
	_ = database.NewUserStore()

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	common.Fatal(err)
	defer l.Close()

	// NewGrpcHandler(grpcServer)

	common.Println("Auth service is started on: ", grpcAddr)
	common.Fatal(grpcServer.Serve(l))
}
