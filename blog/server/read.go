package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog function was invoked with :%v \n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		log.Printf("Invalid argument :%v \n", err)
		return nil, err
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		log.Printf("No document was found with ID :%v \n", in.Id)
		return nil, err
	}

	blog := Marshall(data)

	return blog, nil
}
