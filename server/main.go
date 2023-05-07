package main

import (
	"fmt"
	"net"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	server := grpc.NewServer()
	defer server.Stop()

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	svc := services.NewAccountServer()
	services.RegisterAccountServer(server, svc)
	// Set reflection
	// usage command
	// evans --reflection
	reflection.Register(server)

	fmt.Println("gRPC server start")
	err = server.Serve(listener)

	fmt.Println(err)

	if err != nil {
		panic(err)
	}

}
