package main

import (
	"flag"
	"fmt"
	"grpc-go-poc/pb"
	"grpc-go-poc/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Inside server.go")

	port := flag.Int("port", 0, "server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	laptopServer := service.NewLaptopServer()
	grpcServer := grpc.NewServer()
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
