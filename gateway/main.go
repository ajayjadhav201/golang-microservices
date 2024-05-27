package main

import (
	"net/http"

	"github.com/ajayjadhav201/common"
	"github.com/ajayjadhav201/gateway/auth"
	"github.com/ajayjadhav201/gateway/cache"
	"github.com/go-redis/redis"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr = common.EnvString("GATEWAY_ADDRESS", ":8080")
)

// C:\Program Files\protoc-26.1-win64  //protoc path
func main() {
	aws := auth.NewAwsS3Service()
	//Make Http Connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping the Redis server
	pong, err := rdb.Ping().Result()
	if err != nil {
		common.Println("Error connecting to Redis:", err)
	} else {
		common.Println("Redis connected:", pong)
		defer rdb.Close()
	}
	//
	mux := http.NewServeMux()
	_ = cache.NewCache(rdb) //redis client connection
	RegisterGrpcClient(mux, aws)
	//
	//
	//
	common.Println("Http Server started on: ", httpAddr)
	common.Fatal(http.ListenAndServe(httpAddr, mux))
}
