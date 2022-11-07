package service

import (
	"context"
	"fmt"
	"grpc-go-poc/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server that provides laptop services
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

// CreateLaptop is a Unary RPC to create a new Laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	// Save laptop to store
	time.Sleep(6 * time.Second)

	if ctx.Err() == context.Canceled {
		fmt.Println("request deadline is cancelled")
		return nil, status.Error(codes.Canceled, "request deadline is cancelled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("request deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "request deadline is exceeded")
	}

	laptop := req.GetLaptop()
	fmt.Printf("saved laptop with id: %s", laptop.Id)

	return &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}, nil
}
