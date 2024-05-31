package main

import (
	"golang-microservices/api-gateway/auth"
	"golang-microservices/api-gateway/cache"
	"golang-microservices/api-gateway/products"
	"golang-microservices/api/pb"
	"golang-microservices/common"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr = common.EnvString("GATEWAY_ADDRESS", "localhost:8080")
)

// C:\Program Files\protoc-26.1-win64  //protoc path
func main() {
	err := godotenv.Load(".env")
	common.Panic(err)
	//
	aws := auth.NewAwsS3Service()
	redis, err := RedisClient()
	if err == nil {
		defer redis.Close()
	}

	//Make Http Connection
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	//
	AuthService(r, aws)
	ProductService(r)
	//
	common.Println("Http Server started on: ", httpAddr)
	common.Fatal(r.Run(":8080"))
}

func AuthService(r *gin.Engine, aws *auth.AwsS3Service) {
	// Authentication Client
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	GrpcAuthClient, err := grpc.NewClient("localhost:2001", creds)
	common.Fatal(err)
	GrpcAuthConnection := pb.NewAuthServiceClient(GrpcAuthClient)
	authGroup := r.Group("/api/v2/") // authgroup
	authClient := auth.NewAuthClient(GrpcAuthConnection, aws)
	authClient.RegisterRoutes(authGroup)
}

//
//

func ProductService(r *gin.Engine) {
	// Product Client
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	_, err := grpc.NewClient("localhost:2002", creds)
	common.Fatal(err)
	productsGroup := r.Group("/api/v2/products") //productsgroup
	productClient := products.NewProductClient()
	productClient.RegisterRoutes(productsGroup)
}

//
//

func RedisClient() (cache.Cache, error) {
	//
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping the Redis server
	pong, err := rdb.Ping().Result()
	if err != nil {
		common.Println("Error connecting to Redis:", err)
		return nil, err
	} else {
		common.Println("Redis connected:", pong)
	}
	return cache.NewCache(rdb), nil
}
