package main

import (
	"context"
	"flag"
	"fmt"
	"grpc-go-poc/pb"
	"grpc-go-poc/sample"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Inside client.go")
	serverAddress := flag.String("address", "", "server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	req := pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// Timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, &req)
	if err != nil {
		log.Fatal("Error in creating laptop request: ", err)
	}

	log.Print(" Create laptop response: ", res)
}
