package main

import (
	"context"

	"github.com/ajayjadhav201/common"
	pb "github.com/ajayjadhav201/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGrpcHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (g *grpcHandler) CreateOrder(ctx context.Context, payload *pb.CreateOrderRequest) (*pb.Order, error) {
	common.Println("New order received")

	o := &pb.Order{
		ID:         "42",
		CustomerID: "AJAY",
	}
	return o, nil

}
