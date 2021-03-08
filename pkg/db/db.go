package db

import (
	"context"
	"log"

	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()
)

// InitCollection ...
func InitCollection() *mongo.Collection {
	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.url")))
	if err != nil {
		log.Fatalf("error in 'Create client': %v", err)
	}

	// Create connect
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("error in 'Create connect': %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("error in 'Check the connection': %v", err)
	}

	return client.Database(viper.GetString("mongo.db")).Collection(viper.GetString("mongo.collection"))

}
