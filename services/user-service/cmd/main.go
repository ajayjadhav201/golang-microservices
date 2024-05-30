package main

import (
	"golang-microservices/api"
	"golang-microservices/common"
	"net"
	"user-service/database"
	"user-service/routes"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("USER_SERVICE_ADDRESS", "localhost:2001")
)

func main() {
	//
	err := godotenv.Load(".env")
	common.Panic(err)

	//
	// userStore := database.NewUserStore()
	dynamodb := database.NewDynamoDb() //final

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	common.Fatal(err)
	defer l.Close()
	userservice := routes.NewUserService(dynamodb) //userStore  //final

	api.RegisterAuthServiceServer(grpcServer, userservice) //final

	common.Println("Auth service is started on: ", grpcAddr)
	common.Fatal(grpcServer.Serve(l))
}
