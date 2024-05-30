# .PHONY: generate

hello:
	@echo "welcome to golang-microservices project"

api_gateway:
	go run ./api-gateway/cmd/main.go

user_service:
	go run ./services/user_service/cmd/main.go

product_service:
	go run ./services/product_service/cmd/main.go

order_service:
	go run ./services/order_service/cmd/main.go

payment_service:
	go run ./services/payment_service/cmd/main.go



# generate proto files
gen:
	protoc --go_out=./api \
    --go-grpc_out=./api \
    ./proto/*.proto
