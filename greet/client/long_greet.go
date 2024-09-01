package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Radames"},
		{FirstName: "Emely"},
		{FirstName: "Mérida"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Failed at getting the LongGreet stream: %v", err)
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

	log.Printf("LongGreet Result %s", res.Result)
}
