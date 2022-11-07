package tests

import (
	"fmt"
	"grpc-go-poc/helper"
	"testing"
)

func TestTokenGenerate(t *testing.T) {
	token, err := helper.Generate()
	if err != nil {
		fmt.Println("unable to generate token: ", err)
	}

	fmt.Println("Generated token: ", token)
}
