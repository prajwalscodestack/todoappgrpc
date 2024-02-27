package main

import (
	"log"
	"net"
	pb "todoappgrpc/todo/proto"
	"todoappgrpc/todo/server/db"

	"google.golang.org/grpc"
)

type Server struct {
	pb.TodoServiceServer
}

var addr string = "0.0.0.0:5051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("failed to listen on:", err)
	}
	log.Println("Listening on:", addr)
	db.Init()
	server := grpc.NewServer()
	pb.RegisterTodoServiceServer(server, &Server{})
	if err := server.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
