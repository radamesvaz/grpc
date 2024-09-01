package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) string {
	log.Println("Read blog was invoked")

	blogID := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), blogID)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Printf("Blog Fetched: %s", res)
	return res.Title
}
