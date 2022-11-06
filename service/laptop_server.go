package service

import (
	"context"
	"fmt"
	"grpc-go-poc/pb"
)

// Server that provides laptop services
type LaptopServer struct {
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

// CreateLaptop is a Unary RPC to create a new Laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	// Save laptop to store
	laptop := req.GetLaptop()
	fmt.Printf("saved laptop with id: %s", laptop.Id)

	return &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}, nil
}
