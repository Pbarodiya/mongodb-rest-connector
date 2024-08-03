package repository

import (
	"context"
	"log"
	"mongodbtees/model"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collname string = "students"

func addStudentObject(ctx context.Context, client *mongo.Client, student model.Student, dbname string) (*mongo.InsertOneResult, error) {
	collection := client.Database(dbname).Collection(collname)
	c, error := collection.InsertOne(ctx, student)
	if error != nil {
		return nil, error
	}
	return c, nil
}

func CreateStudentObject(ctx context.Context, client *mongo.Client, dbname string) {
	stud1 := model.Student{
		Name: "John Doe",
		Age:  20,
		Grade: model.Grade{
			Subject: "Mathematics",
			Score:   85,
		}}

	stud2 := model.Student{
		Name: "Mohn Doe",
		Age:  22,
		Grade: model.Grade{
			Subject: "Mathematics",
			Score:   80,
		}}

	stud3 := model.Student{
		Name: "Sohn Doe",
		Age:  21,
		Grade: model.Grade{
			Subject: "Mathematics",
			Score:   20,
		}}

	stud4 := model.Student{
		Name: "Rohn Doe",
		Age:  21,
		Grade: model.Grade{
			Subject: "Mathematics",
			Score:   72,
		}}

	stud5 := model.Student{
		Name: "Pohn Doe",
		Age:  23,
		Grade: model.Grade{
			Subject: "Mathematics",
			Score:   10,
		}}

	var ListOfStudent = []model.Student{stud1, stud2, stud3, stud4, stud5}

	for _, L := range ListOfStudent {
		_, error := addStudentObject(ctx, client, L, dbname)
		if error != nil {
			log.Fatal(error)
		}
	}
}

func FetchStudentByID(ctx context.Context, client *mongo.Client, ID string, dbname string) (*model.Student, error) {
	collection := client.Database(dbname).Collection(collname)

	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	// filterResult, err := collection.Find(ctx, bson.M{})
	// fmt.Println("FIlter result : ", filterResult)
	// fmt.Println("error : ", err)

	filter := bson.M{"_id": objectID}

	studentBson := collection.FindOne(ctx, filter)

	var studentDocument *model.Student
	err = studentBson.Decode(&studentDocument)
	if err != nil {
		return nil, err
	}

	return studentDocument, nil

}
