package main

import (
	"net/http"

	"github.com/ajayjadhav201/common"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr         = common.EnvString("GATEWAY_ADDRESS", ":8080")
	orderServiceAddr = "localhost:2000"
	authServiceAddr  = "localhost:2001"
)

// C:\Program Files\protoc-26.1-win64  //protoc path
func main() {
	//Make Http Connection
	mux := http.NewServeMux()
	RegisterGrpcClient(mux)
	//
	//
	//
	common.Println("Http Server started on: ", httpAddr)
	common.Fatal(http.ListenAndServe(httpAddr, mux))
}
