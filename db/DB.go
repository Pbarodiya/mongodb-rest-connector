// db/db.go
package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017" // Replace with your MongoDB connection string
const dbName = "school"

var mongoClient *mongo.Client

func GetMongoClient() (*mongo.Client, context.Context, string) {
	if mongoClient != nil {
		return mongoClient, context.TODO(), dbName
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	mongoClient = client

	return mongoClient, context.TODO(), dbName
}
