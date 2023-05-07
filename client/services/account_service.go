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
	Average(numbers ...float64) error
	Sum(numbers ...float64) error
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
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

func (c *accountService) Average(numbers ...float64) error {
	stream, err := c.accountClient.Average(context.Background())

	if err != nil {
		return err
	}

	for _, number := range numbers {

		req := AverageRequest{
			Number: number,
		}

		err := stream.Send(&req)

		if err != nil {
			return nil
		}

		fmt.Printf("Streaming data sent: %v \n", number)

		time.Sleep(time.Second * 1)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		return err
	}

	fmt.Printf("Response: %v \n", res.Value)

	return nil
}

func (c *accountService) Sum(numbers ...float64) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	stream, err := c.accountClient.Sum(ctx)
	quit := make(chan bool)
	errCh := make(chan error)

	if err != nil {
		return err
	}

	go func() {
		for _, number := range numbers {

			req := SumRequest{
				Number: number,
			}

			stream.Send(&req)

			fmt.Printf("Request sent: %v \n", number)

			time.Sleep(time.Second * 1) // Delay
		}

		fmt.Println("Close Send")
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				quit <- true
				return
			} else if err != nil {
				errCh <- err
				return
			}

			fmt.Printf("Response (sum): %v \n", res.Result)
			time.Sleep(time.Second * 1)
		}
	}()

	fmt.Println("Wait...")
	select {
	case <-quit:
		return nil
	case err := <-errCh:
		fmt.Printf("Error: %v \n", err)
		return err
	case <-ctx.Done():
		fmt.Println(ctx.Err().Error())
		return ctx.Err()
	}
}
