package routers

import (
	"mongodbtees/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	app.Get("/students/:id", handler.CalculateReport)
}
