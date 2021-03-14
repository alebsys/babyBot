package db

import (
	"context"
	l "github.com/alebsys/baby-bot/pkg/logs"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()
)

// InitCollection ...
func InitCollection() *mongo.Collection {
	l.Sugar.Debugf("Trying to init mongo collection: '%v'", viper.GetString("mongo.collection"))
	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.url")))
	if err != nil {
		l.Sugar.Fatal(err)
	}

	// Create connect
	err = client.Connect(ctx)
	if err != nil {
		l.Sugar.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		l.Sugar.Fatal(err)
	}

	l.Sugar.Debugf("Mongo collection has been initialized: '%v'", viper.GetString("mongo.collection"))
	return client.Database(viper.GetString("mongo.db")).Collection(viper.GetString("mongo.collection"))

}
