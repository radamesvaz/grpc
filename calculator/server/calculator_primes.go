package main

import (
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) CalculatorPrimes(in *pb.CalculatorRequest, stream pb.CalculatorService_CalculatorPrimesServer) error {
	log.Printf("CalculatorPrimes function was invoked with %v", in)

	k := uint64(2)
	n := in.FirstNumber

	for n > 1 {
		if n%k == 0 {

			stream.Send(&pb.CalculatorResponse{
				Result: k,
			})

			n = n / k
		} else {
			k = k + 1
		}
	}

	return nil
}
