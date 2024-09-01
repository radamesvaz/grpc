package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) string {
	log.Println("Update blog was invoked")

	blog := &pb.Blog{
		Id:       id,
		AuthorId: "Radames",
		Title:    "This blog has been updated",
		Content:  "Without waiting for the course",
	}

	_, err := c.UpdateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Printf("Blog with ID: %v has been updated:", id)
	return id
}
