package services

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type accountServer struct {
}

func NewAccountServer() AccountServer {
	return &accountServer{}
}

func (accountServer) User(ctx context.Context, req *UserRequest) (*UserResponse, error) {

	response := UserResponse{
		Result:      "Name: " + req.Name,
		CreatedDate: timestamppb.Now(),
	}

	return &response, nil
}

func (accountServer) Fibonacci(req *FibonacciRequest, stream Account_FibonacciServer) error {

	for n := uint32(0); n <= req.N; n++ {
		result := fib(n)

		res := FibonacciResponse{
			Result: result,
		}

		// Streaming
		err := stream.Send(&res)

		if err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	// md := make(metadata.MD)

	// md["A"] = []string{"a", "b", "C"}

	// stream.SetHeader(md)

	return nil
}

func fib(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func (accountServer) mustEmbedUnimplementedAccountServer() {}
