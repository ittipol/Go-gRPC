package services

import (
	"context"
	"fmt"
)

type AccountService interface {
	User(name string) error
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
