package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:5051"

const mongoConn = "mongodb://root:root@localhost:27017/"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoConn))
	if err != nil {
		log.Fatalf("Failed to setup the mongo db client: %v", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v. Error: %v", addr, err)
	}

	log.Printf("Listening on port: %v", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
