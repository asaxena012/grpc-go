package service_test

import (
	"context"
	"fmt"
	"grpc-go-poc/pb"
	"grpc-go-poc/sample"
	"grpc-go-poc/service"
	"net"
	"testing"

	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	_, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		fmt.Println("Error in creating laptop request", err)
	}
	fmt.Println(" Create laptop response: ", res)

}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer()

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error in starting server", err)
	}

	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Error in connecting client to target: ", err)
	}

	return pb.NewLaptopServiceClient(conn)
}
