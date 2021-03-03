package db

import (
	"context"
	"github.com/spf13/viper"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()
)

func InitCollection() *mongo.Collection {
	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.url")))
	if err != nil {
		log.Fatalf("Error in 'Create client': %v", err)
	}

	// Create connect
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error in 'Create connect': %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error in 'Check the connection': %v", err)
	}

	return client.Database(viper.GetString("mongo.db")).Collection(viper.GetString("mongo.collection"))

}
