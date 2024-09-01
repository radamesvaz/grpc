package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog function was invoked with :%v \n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		log.Printf("Invalid argument :%v \n", err)
		return nil, err
	}

	updatedBlog := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	filter := bson.M{"_id": oid}
	data := &BlogItem{}

	res := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": updatedBlog})
	if err := res.Decode(data); err != nil {
		log.Printf("No document was found with ID :%v \n", in.Id)
		return nil, err
	}

	blog := Marshall(data)
	log.Printf("Blog was updated successfully :%v \n", blog)

	return nil, nil
}
