package sample

import (
	"grpc-go-poc/pb"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
)

func NewLaptop() *pb.Laptop {

	keyboard := pb.Keyboard{
		Layout:  pb.Keyboard_QWERTY,
		Backlit: true,
	}

	cpu := pb.CPU{
		Brand:   "intel",
		Name:    "I9",
		Threads: 16,
		MinGhz:  1.2,
		MaxGhz:  5.6,
	}

	gpu := pb.GPU{
		Brand:  "intel",
		Name:   "I9",
		MinGhz: 1.2,
		MaxGhz: 5.6,
		Memory: &pb.Memory{
			Value: 2,
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	storage := pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: 1,
			Unit:  pb.Memory_TERABYTE,
		},
	}

	screen := pb.Screen{
		SizeInch: 15.6,
		Resolution: &pb.Screen_Resolution{
			Width:  9,
			Height: 5,
		},
		Panel:      pb.Screen_IPS,
		Multitouch: false,
	}

	laptop := pb.Laptop{
		Id:    uuid.New().String(),
		Brand: "Alienware",
		Name:  "Phantom",
		Cpu:   &cpu,
		Ram: &pb.Memory{
			Value: 16,
			Unit:  pb.Memory_GIGABYTE,
		},
		Gpus:     []*pb.GPU{&gpu},
		Storages: []*pb.Storage{&storage},
		Screen:   &screen,
		Keyboard: &keyboard,

		Weight: &pb.Laptop_WeightKg{
			WeightKg: 1.2,
		},

		PriceUsd:    400,
		ReleaseYear: 2018,
		UpdatedAt:   &timestamp.Timestamp{},
	}

	return &laptop
}
