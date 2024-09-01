package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) CalculatorAvg(stream pb.CalculatorService_CalculatorAvgServer) error {
	log.Print("CalculatorAvg was invoked")

	res := float32(0)
	reqNumber := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			res = res / float32(reqNumber)
			return stream.SendAndClose(&pb.CalculatorAvgResponse{
				Result: float32(res),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading the client stream: %v", err)
		}

		res += float32(req.GetFirstNumber())
		reqNumber++
	}

}
