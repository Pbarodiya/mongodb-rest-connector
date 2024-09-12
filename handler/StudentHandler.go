package handler

import (
	"fmt"
	"log"
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
	totalmarks := 0
	for _, grade := range student.Grade {
		totalmarks += grade.Score
	}
	percentage := (totalmarks) / len(student.Grade)

	var report1 = model.Report{
		ID:         student.ID,
		Name:       student.Name,
		Age:        student.Age,
		TotalMarks: totalmarks,
		Percentage: float32(percentage),
	}
	client, ctx, _ := db.GetMongoClient()
	_, err := repository.StoreReport(ctx, client, "school", &report1)
		if err != nil {
			log.Fatal(err)
		}

		return report1

	}

