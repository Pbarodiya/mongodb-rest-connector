package main

import (
	"context"
	"fmt"

	//"fmt"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"mongodbtees/repository"
	"mongodbtees/routers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017" // Replace with your MongoDB connection string
const dbName = "school"

var mongoClient *mongo.Client

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	mongoClient = client
	repository.CreateStudentObject(context.TODO(), mongoClient, dbName)

	app := fiber.New()
	routers.SetupRouters(app)
	fmt.Println("server is running on port 3000")
	log.Fatal(app.Listen(":3000"))
}

// func getMongoClient() (*mongo.Client, context.Context, string) {
// 	if mongoClient != nil {
// 		return mongoClient, context.TODO(), dbName
// 	}
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	mongoClient = client

// 	return mongoClient, context.TODO(), dbName

// }
