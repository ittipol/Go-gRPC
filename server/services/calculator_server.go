package services

import "context"

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return &calculatorServer{}
}

func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

	response := HelloResponse{
		Result: "Req: " + req.Name,
	}

	return &response, nil
}

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}
