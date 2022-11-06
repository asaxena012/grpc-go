package service

import "grpc-go-poc/pb"

type LaptopStore interface {
	Save(*pb.Laptop) error
}
