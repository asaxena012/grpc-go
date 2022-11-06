package serializer_test

import (
	"fmt"
	"grpc-go-poc/pb"
	"grpc-go-poc/sample"
	"grpc-go-poc/serializer"
	"testing"
)

func TestSerializerWrite(t *testing.T) {
	laptop := sample.NewLaptop()

	binaryFile := "../tmp/laptop.bin"

	err := serializer.WritePbToBinaryFile(laptop, binaryFile)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func TestSerializerRead(t *testing.T) {
	laptop := &pb.Laptop{}

	binaryFile := "../tmp/laptop.bin"

	err := serializer.ReadPbFromBinaryFile(laptop, binaryFile)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(laptop)
}

func TestPbToJson(t *testing.T) {
	laptop := sample.NewLaptop()

	data, err := serializer.PbToJson(laptop)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(data)
}
