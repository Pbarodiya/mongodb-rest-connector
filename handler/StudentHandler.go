package handler

import (
	"fmt"
	"mongodbtees/db"
	"mongodbtees/model"
	"mongodbtees/repository"

	"github.com/gofiber/fiber/v2"
)

func CalculateReport(c *fiber.Ctx) error {
	id := c.Params("id")

	client, ctx, dbName := db.GetMongoClient()

	studDocument, err := repository.FetchStudentByID(ctx, client, id, dbName)
	if err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"message": err})
	}

	reportOfStudent := Calculate(studDocument)
	return c.Status(200).JSON(reportOfStudent)

}

func Calculate(student *model.Student) model.Report {
	totalmarks := student.Grade.Score
	percentage := totalmarks * 100 / 100

	return model.Report{
		ID:         student.ID,
		Name:       student.Name,
		Age:        student.Age,
		TotalMarks: totalmarks,
		Percentage: float32(percentage),
	}
}
