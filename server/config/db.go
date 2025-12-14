package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TodoCollection *mongo.Collection


func ConnectDB(){
	client , err := mongo.Connect(context.Background() , options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	TodoCollection = client.Database("GO_TODO").Collection("todos")
	log.Println("MongoDB Connected")
}