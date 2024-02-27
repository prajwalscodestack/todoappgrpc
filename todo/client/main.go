package main

import (
	"context"
	"log"
	pb "todoappgrpc/todo/proto"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5051"

type Client struct {
}

func saveTodo(client pb.TodoServiceClient) {
	log.Println("Client is calling the server")
	res, err := client.SaveTodo(context.Background(), &pb.TodoRequest{
		Id:          ksuid.New().String(),
		Title:       "Create Todo App based on GRPC",
		Desc:        "Todo app should have server and client and todos should be saved in mongodb",
		IsCompleted: true,
	})
	if err != nil {
		log.Fatal("error while calling method:", err)
	}
	log.Println("response:", res.Result)
}
func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect to server", err)
	}
	client := pb.NewTodoServiceClient(conn)
	saveTodo(client)
}
