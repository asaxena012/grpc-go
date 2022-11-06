package service_test

import (
	"context"
	"grpc-go-poc/pb"
	"grpc-go-poc/sample"
	"grpc-go-poc/service"
	"testing"
)

func TestServerCreateLaptop(t *testing.T) {

	laptop := sample.NewLaptop()

	server := service.NewLaptopServer()

	req := pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	server.CreateLaptop(context.Background(), &req)

}
