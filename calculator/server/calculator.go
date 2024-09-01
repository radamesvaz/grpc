package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v", in)
	sum := in.FirstNumber + in.SecondNumber
	return &pb.CalculatorResponse{
		Result: sum,
	}, nil
}
