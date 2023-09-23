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

	app.Post("/api/step/:id", middlewares.AuthMiddleware(), stepController.CreateStep)
	app.Get("/api/step", stepController.GetAllStep)
	app.Get("/api/step/:id", stepController.GetStepById)
	app.Get("/api/step/menu/:id", stepController.GetStepByMenuId)
	app.Put("/api/step/:id", middlewares.AuthMiddleware(), stepController.UpdateStep)
	app.Delete("/api/step/:id", middlewares.AuthMiddleware(), stepController.DeleteStepById)
}
