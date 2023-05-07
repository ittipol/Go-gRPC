package services

import (
	"context"
	"fmt"
	"io"
	"time"
)

type AccountService interface {
	User(name string) error
	Fibonacci(n uint32) error
}

type accountService struct {
	accountClient AccountClient
}

func NewAccountService(accountClient AccountClient) AccountService {
	return &accountService{accountClient}
}

func (c *accountService) User(name string) error {
	req := UserRequest{
		Name: name,
	}

	response, err := c.accountClient.User(context.Background(), &req)

	if err != nil {
		return err
	}

	fmt.Printf("Response: %#v \n", response)
	fmt.Printf("Response: %#v \n", response.Result)

	return nil
}

func (c *accountService) Fibonacci(n uint32) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	req := FibonacciRequest{
		N: n,
	}

	stream, err := c.accountClient.Fibonacci(ctx, &req)

	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		fmt.Printf("Fibonacci Response: %v \n", res.Result)
	}

	return nil
}
