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
	"google.golang.org/grpc/metadata"
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

	md := metadata.New(map[string]string{
		"Authorization": "shdlakdk3h4ui23h4iu",
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := laptopClient.CreateLaptop(ctx, &req)
	if err != nil {
		log.Fatal("Error in creating laptop request: ", err)
	}

	log.Print(" Create laptop response: ", res)
}
