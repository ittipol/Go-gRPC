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

		// fmt.Printf("CERT: %v \n", cert)

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

	var selected string
	quit := false
	for {
		fmt.Println("1. Test gRPC \"Unary\"")
		fmt.Println("2. Test gRPC \"Server Streaming\"")
		fmt.Println("3. Test gRPC \"Client Streaming\"")
		fmt.Println("4. Test gRPC \"Bidirectional Streaming\"")
		fmt.Println("Input [1, 2, 3, 4]:")
		fmt.Scanf("%s", &selected)

		fmt.Printf("selected: %v \n", selected)

		switch selected {
		case "1":

			var name string
			fmt.Println("Input name:")
			fmt.Scanf("%s", &name)

			err = accountService.User(name)
		case "2":
			err = accountService.Fibonacci(4)
		case "3":
			err = accountService.Average(3, 4, 5, 6, 7, 8, 9, 20, 100, 10000, 1234)
		case "4":
			err = accountService.Sum(3, 4, 5, 6, 7, 8, 9, 20, 100, 10000, 9876)
		default:
			quit = true
		}

		if err != nil || quit {
			fmt.Println("Exit")
			break
		}

		fmt.Printf("\n\n")
	}

	if err != nil {

		if grpcError, ok := status.FromError(err); ok {
			fmt.Printf("gRPC occured an error\n")
			fmt.Printf("%v | %v | %v \n", grpcError.Code(), grpcError.Details(), grpcError.Message())
			panic(grpcError)
		}

		panic(err)
	}
}
