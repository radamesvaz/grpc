package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient) {
	log.Println("doCalculator was invoked")
	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		FirstNumber:  28,
		SecondNumber: 78,
	})

	if err != nil {
		log.Fatalf("Error with CalculatorRequest: %v", err)
	}

	log.Printf("Calculatoring: %v", res.Result)
}
