package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"todoappgrpc/todo/server/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var todoDAO *mongo.Database

func Init() {
	connectionString := "mongodb://localhost:27017"

	// Create a MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	todoDAO = client.Database("grpc")
}

func SaveTodo(todo models.Todo) error {
	_, err := todoDAO.Collection("todos").InsertOne(context.TODO(), todo)
	if err != nil {
		fmt.Print("err:", err)
		return err
	}
	return nil
}
