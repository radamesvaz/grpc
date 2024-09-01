package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Radamés",
	})

	if err != nil {
		log.Fatalf("Error with GreetRequest: %v", err)
	}

	log.Printf("Greeting: %s", res.Result)
}
