package main

import (
	"context"
	"flag"
	"fmt"
	"grpc-go-poc/pb"
	"grpc-go-poc/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("unary interceptor called: ", info.FullMethod)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Failed to get metadata")
	}

	token := md.Get("Authorization")[0]

	log.Println("Authorization token in request: ", token)

	// Parse and validate token

	return handler(ctx, req)
}

func main() {
	fmt.Println("Inside server.go")

	port := flag.Int("port", 0, "server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	laptopServer := service.NewLaptopServer()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Println("Error in creating tcp listener: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println("Unable to start server: ", err)
	}
}
