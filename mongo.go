package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongo() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	uri := os.Getenv("ATLAS_MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

var DB *mongo.Client = InitMongo()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.
		Database("cluster0").
		Collection(collectionName)

	return collection
}
