package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//collection *mongo.Collection
	ctx = context.TODO()
)

func InitCollection() *mongo.Collection {
	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		fmt.Println("error in 'Create client'")
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("error in 'Create connect'")
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("error in 'Check the connection'")
		log.Fatal(err)
	}

	return client.Database("test").Collection("trainers")

}
