# .PHONY: generate

hello:
	@echo "welcome to golang-microservices project"


# api-gateway commands
api_gateway:
	go run ./api-gateway/cmd/main.go
#  go get ${P}

# user-service commands
user_service:
	go run ./services/user-service/cmd/main.go

# product-service commands
product_service:
	go run ./services/product-service/cmd/main.go

# order-service commands
order_service:
	go run ./services/order-service/cmd/main.go

# payment-service commands
payment_service:
	go run ./services/payment-service/cmd/main.go



# generate proto files
gen:
	protoc --go_out=./api \
    --go-grpc_out=./api \
    ./proto/*.proto
