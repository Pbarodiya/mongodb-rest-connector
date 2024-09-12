package main

import (
	"context"
	"fmt"

	"log"
	"mongodbtees/repository"
	"mongodbtees/routers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017" // Replace with your MongoDB connection string
//process which serves on HTTP, always has host and port, one which serves on personal machine has localhost 
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
	count, err := mongoClient.Database(dbName).Collection("students").EstimatedDocumentCount(context.TODO())
	if err != nil {
		panic(err)
	 }

	 // if students document has more than 5 docs this function should not run
	if count < 5 {
	 
	repository.CreateStudentObject(context.TODO(), mongoClient, dbName)

	}
	app := fiber.New()        //sets up the Fiber app
	routers.SetupRouters(app) //configures the HTTP routes using SetupRouters
	fmt.Println("server is running on port 3000")
	log.Fatal(app.Listen(":3000"))
}