package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Age   int                `bson:"age"`
	Grade Grade              `bson:"grade"`
}

type Grade struct {
	Subject string `bson:"subject"`
	Score   int    `bson:"score"`
}

type Report struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Age        int                `bson:"age"`
	TotalMarks int                `bson:"total_marks"`
	Percentage float32            `bson:"percentage"`
}
