package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func calculatorPrimes(c pb.CalculatorServiceClient) {
	log.Println("calculatorPrimes was invoked")

	req := &pb.CalculatorRequest{
		FirstNumber: 120,
	}

	stream, err := c.CalculatorPrimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling calculatorPrimes %v", err)
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

		log.Printf("calculatorPrimes: %v", msg.Result)
	}

}
