package main

import (
	"context"
	"log"
	pb "todoappgrpc/todo/proto"
	"todoappgrpc/todo/server/db"
	"todoappgrpc/todo/server/models"
)

func (s *Server) SaveTodo(ctx context.Context, in *pb.TodoRequest) (*pb.TodoResponse, error) {
	log.Println("SaveTodo Function Invoked!")
	todo := &models.Todo{
		Id:          in.Id,
		Title:       in.Title,
		Desc:        in.Desc,
		IsCompleted: in.IsCompleted,
	}
	if err := db.SaveTodo(*todo); err != nil {
		return &pb.TodoResponse{
			Result: "Failed to store todo:",
		}, err
	}
	return &pb.TodoResponse{
		Result: "Todo Saved Successfully!",
	}, nil
}
