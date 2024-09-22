package main

import (
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect on: %v, Error: %v", addr, err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	// create blog
	blogID := createBlog(c)

	// read blog
	readBlog(c, blogID)

	// update blog
	updateBlog(c, blogID)

	// list blog
	listBlog(c)
}
