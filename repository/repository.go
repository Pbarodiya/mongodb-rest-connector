package repository

import (
	"context"
	"log"
	"mongodbtees/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collname string = "students"
var reportscoll string = "report"

func addStudentObject(ctx context.Context, client *mongo.Client, student model.Student, dbname string) (*mongo.InsertOneResult, error) {
	collection := client.Database(dbname).Collection(collname)
	c, error := collection.InsertOne(ctx, student)
	if error != nil {
		return nil, error
	}
	return c, nil
}

// The CreateStudentObject function creates multiple student objects and inserts them into the MongoDB collection specified by collname (which is "students" )
func CreateStudentObject(ctx context.Context, client *mongo.Client, dbname string) {
	stud1 := model.Student{
		Name: "John Doe",
		Age:  20,
		Grade: []model.Grade{
			{
				Subject: "Mathematics",
				Score:   85,
			},
			{
				Subject: "English",
				Score:   80,
			},
			{
				Subject: "Science",
				Score:   72,
			},
		},
	}

	stud2 := model.Student{
		Name: "Dohn Doe",
		Age:  21,
		Grade: []model.Grade{
			{
				Subject: "Mathematics",
				Score:   48,
			},
			{
				Subject: "English",
				Score:   65,
			},
			{
				Subject: "Science",
				Score:   68,
			},
		},
	}

	stud3 := model.Student{
		Name: "Rianna Leefar",
		Age:  22,
		Grade: []model.Grade{
			{
				Subject: "Mathematics",
				Score:   66,
			},
			{
				Subject: "English",
				Score:   72,
			},
			{
				Subject: "Science",
				Score:   91,
			},
		},
	}

	stud4 := model.Student{
		Name: "Mark leal",
		Age:  22,
		Grade: []model.Grade{
			{
				Subject: "Mathematics",
				Score:   38,
			},
			{
				Subject: "English",
				Score:   55,
			},
			{
				Subject: "Science",
				Score:   45,
			},
		},
	}

	stud5 := model.Student{
		Name: "Dario Leal",
		Age:  23,
		Grade: []model.Grade{
			{
				Subject: "Mathematics",
				Score:   89,
			},
			{
				Subject: "English",
				Score:   87,
			},
			{
				Subject: "Science",
				Score:   90,
			},
		},
	}

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

	filter := bson.M{"_id": objectID}

	studentBson := collection.FindOne(ctx, filter)

	var studentDocument *model.Student
	err = studentBson.Decode(&studentDocument)
	if err != nil {
		return nil, err
	}

	return studentDocument, nil

}

func StoreReport(ctx context.Context, client *mongo.Client, dbname string, report *model.Report) (*mongo.InsertOneResult, error) {
	collection := client.Database(dbname).Collection(reportscoll)
	r, error := collection.InsertOne(ctx, report)
	if error != nil {
		return nil, error
	}
	return r, nil
}
