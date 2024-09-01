package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doCalculatorAvg(c pb.CalculatorServiceClient) {
	log.Println("doCalculatorAvg was invoked")

	reqs := []*pb.CalculatorRequest{
		{FirstNumber: 1},
		{FirstNumber: 2},
		{FirstNumber: 3},
		{FirstNumber: 4},
	}

	stream, err := c.CalculatorAvg(context.Background())

	if err != nil {
		log.Fatalf("Error with CalculatorRequest: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		// Adding a sleep to make the logs more readable
		// Doesn't have anything to do with how the stream works
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while sending the request to the server: %v", err)
	}

	log.Printf("doCalculatorAvg Result %v", res.Result)
}
