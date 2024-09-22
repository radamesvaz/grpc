package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("LIST blog was invoked")

	stream, err := c.ListBlogs(context.Background(), &empty.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}

		log.Println(res)
	}
}
