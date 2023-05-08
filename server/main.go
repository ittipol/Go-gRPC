package main

import (
	"flag"
	"fmt"
	"net"
	"server/services"
	"strings"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {

	var server *grpc.Server

	// Receive value from Flag in command-line
	tls := flag.Bool("tls", false, "Use TLS Connection") // Create Flag name tls
	flag.Parse()
	// See list of flag after parsed
	// go run . -h

	if flag.Parsed() && *tls {
		cert := viper.GetString("tls.cert")
		key := viper.GetString("tls.key")

		fmt.Printf("CERT: %v | KEY: %v \n", cert, key)

		creds, err := credentials.NewServerTLSFromFile(cert, key)

		if err != nil {
			panic(err)
		}

		server = grpc.NewServer(grpc.Creds(creds))

		fmt.Println("Connection is secure")
	} else {
		// insecure
		server = grpc.NewServer()

		fmt.Println("Connection is not secure")
	}

	defer server.Stop()

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	svc := services.NewAccountServer()
	services.RegisterAccountServer(server, svc)
	// Set reflection
	// Usage command
	// evans --reflection
	reflection.Register(server)

	fmt.Println("gRPC server start")
	err = server.Serve(listener)

	fmt.Println(err)

	if err != nil {
		panic(err)
	}

}
