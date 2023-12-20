package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


var client *mongo.Client

func init() {
	client = connect()
}

func connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		close(ctx)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("failed to connect to mongodb", err)
	}

	fmt.Println("Successfully connected to mongodb")

	return client
}

func close(ctx context.Context) {
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

const dbName = "restaurant"

func GetCollection(collName string) *mongo.Collection {
	coll := client.Database(dbName).Collection(collName)
	return coll
}
