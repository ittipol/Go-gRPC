package main

import (
	"fmt"
	"net"
	"server/services"

	"google.golang.org/grpc"
)

func main() {
	// Start gRPC Server
	server := grpc.NewServer()
	defer server.Stop()

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	srv := services.NewAccountServer()
	services.RegisterAccountServer(server, srv)

	fmt.Println("gRPC server start")
	err = server.Serve(listener)

	fmt.Println(err)

	if err != nil {
		panic(err)
	}

}
