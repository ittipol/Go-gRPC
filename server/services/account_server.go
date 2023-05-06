package services

import "context"

type accountServer struct {
}

func NewAccountServer() AccountServer {
	return &accountServer{}
}

func (accountServer) User(ctx context.Context, req *UserRequest) (*UserResponse, error) {

	response := UserResponse{
		Result: "Name: " + req.Name,
	}

	return &response, nil
}

func (accountServer) mustEmbedUnimplementedAccountServer() {}
