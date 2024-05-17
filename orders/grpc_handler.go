package main

import (
	"api"
	"context"

	"github.com/ajayjadhav201/common"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	api.UnimplementedOrderServiceServer
}

func NewGrpcHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	api.RegisterOrderServiceServer(grpcServer, handler)
}

func (g *grpcHandler) CreateOrder(ctx context.Context, payload *api.CreateOrderRequest) (*api.Order, error) {
	common.Println("New order received")

	o := &api.Order{
		ID:         "42",
		CustomerID: "AJAY",
	}
	return o, nil

}
