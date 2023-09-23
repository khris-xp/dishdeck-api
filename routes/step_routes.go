package routes

import (
	"dishdeck-api/controllers"
	"dishdeck-api/middlewares"
	"dishdeck-api/repositories"

	"github.com/gofiber/fiber/v2"
)

func StepRoutes(app *fiber.App) {
	stepRepo := repositories.NewStepRepository()
	stepController := controllers.NewStepController(stepRepo)

	app.Post("/api/step", middlewares.AuthMiddleware(), stepController.CreateStep)
	app.Get("/api/step", stepController.GetAllStep)
}
