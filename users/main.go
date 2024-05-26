package main

import (
	"api"
	"net"

	"github.com/ajayjadhav201/common"
	"github.com/ajayjadhav201/users/database"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("USER_SERVICE_ADDRESS", "localhost:2001")
)

func main() {
	// err := godotenv.Load(".env")
	// common.Panic(err)

	//
	_ = database.NewUserStore()

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	common.Fatal(err)
	defer l.Close()
	userservice := NewUserService()

	api.RegisterAuthServiceServer(grpcServer, userservice)

	common.Println("Auth service is started on: ", grpcAddr)
	common.Fatal(grpcServer.Serve(l))
}
