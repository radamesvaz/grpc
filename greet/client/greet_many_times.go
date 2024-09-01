package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Radamés",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes %v", err)
	}

	for {
		msg, err := stream.Recv()

		/*
		* EOF is the error returned by Read when no more input is available.
		* (Read must return EOF itself, not an error wrapping EOF, because callers will test for EOF using ==.)
		* Functions should return EOF only to signal a graceful end of input.
		* If the EOF occurs unexpectedly in a structured data stream, the appropriate error is either ErrUnexpectedEOF or some other error giving more detail.
		 */
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v", err)
		}

		log.Printf("GreetManyTimes: %s", msg.GetResult())
	}

}
