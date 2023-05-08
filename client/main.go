package main

import (
	"client/services"
	"flag"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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

	var (
		cc    *grpc.ClientConn
		err   error
		creds credentials.TransportCredentials
	)

	host := flag.String("host", "localhost:50051", "gRPC Server host")
	tls := flag.Bool("tls", false, "Use TLS Connection")
	flag.Parse()

	creds = insecure.NewCredentials()

	if flag.Parsed() && *tls {
		cert := viper.GetString("tls.cert")

		fmt.Printf("CERT: %v \n", cert)

		creds, err = credentials.NewClientTLSFromFile(cert, "")

		if err != nil {
			panic(err)
		}
	}

	cc, err = grpc.Dial(*host, grpc.WithTransportCredentials(creds))

	if err != nil {
		panic(err)
	}

	defer cc.Close()

	accountClient := services.NewAccountClient(cc)
	accountService := services.NewAccountService(accountClient)

	// err = accountService.User("Hello")
	// err = accountService.Fibonacci(4)
	err = accountService.Average(3, 4, 5, 6, 7, 8, 9, 20, 100, 10000)
	// err = accountService.Sum(3, 4, 5, 6, 7, 8, 9, 20, 100, 10000)

	if err != nil {

		if grpcError, ok := status.FromError(err); ok {
			fmt.Printf("gRPC occured an error\n")
			fmt.Printf("%v | %v | %v \n", grpcError.Code(), grpcError.Details(), grpcError.Message())
			panic(grpcError)
		}

		panic(err)
	}
}
