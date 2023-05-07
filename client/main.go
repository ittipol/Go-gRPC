package main

import (
	"client/services"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {

	cred := insecure.NewCredentials()

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(cred))
	defer cc.Close()

	if err != nil {
		panic(err)
	}

	accountClient := services.NewAccountClient(cc)
	accountService := services.NewAccountService(accountClient)

	err = accountService.User("")
	// err = accountService.Fibonacci(4)
	// err = accountService.Average(3, 4, 5, 6, 7, 8, 9, 20, 100, 10000)
	// err = accountService.Sum(3, 4, 5, 6, 7, 8, 9, 20, 100, 10000)

	if err != nil {

		if grpcError, ok := status.FromError(err); ok {
			fmt.Printf("Is gRPC Error \n")
			fmt.Printf("%v | %v | %v \n", grpcError.Code(), grpcError.Details(), grpcError.Message())
			panic(grpcError)
		}

		panic(err)
	}
}
