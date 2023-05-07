package main

import (
	"client/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	// err = accountService.User("New User")
	err = accountService.Fibonacci(5)

	if err != nil {
		panic(err)
	}
}
