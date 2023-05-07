package services

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type accountServer struct {
}

func NewAccountServer() AccountServer {
	return &accountServer{}
}

func (accountServer) User(ctx context.Context, req *UserRequest) (*UserResponse, error) {

	if req.Name == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Name is required",
		)
	}

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

func (accountServer) Average(stream Account_AverageServer) error {

	var (
		sum, avg float64
		count    int32
	)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil
		}

		fmt.Printf("Request (Stream): %v \n", req.Number)

		sum += req.Number
		count++
	}

	avg = sum / float64(count)
	fmt.Printf("Total: %v \n", sum)
	fmt.Printf("Avg: %v \n", avg)

	res := AvereageResponse{
		Value: avg,
	}

	return stream.SendAndClose(&res)
}

func (accountServer) Sum(stream Account_SumServer) error {

	var sum float64

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		sum += req.Number

		fmt.Printf("Request (Stream): %v \n", req.Number)

		res := SumResponse{
			Result: sum,
		}

		// Send to client
		err = stream.Send(&res)

		if err != nil {
			return err
		}
	}

	println("Exit...")
	return nil
}

func (accountServer) mustEmbedUnimplementedAccountServer() {}
