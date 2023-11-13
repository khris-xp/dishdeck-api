package routes

import (
	"dishdeck-api/controllers"
	"dishdeck-api/middlewares"
	"dishdeck-api/repositories"

	"github.com/gofiber/fiber/v2"
)

const StepAPIPath = "/api/step/:id"

func StepRoutes(app *fiber.App) {
	stepRepo := repositories.NewStepRepository()
	stepController := controllers.NewStepController(stepRepo)

	app.Post(StepAPIPath, middlewares.AuthMiddleware(), stepController.CreateStep)
	app.Get("/api/step", stepController.GetAllStep)
	app.Get(StepAPIPath, stepController.GetStepById)
	app.Get("/api/step/menu/:id", stepController.GetStepByMenuId)
	app.Put(StepAPIPath, middlewares.AuthMiddleware(), stepController.UpdateStep)
	app.Delete(StepAPIPath, middlewares.AuthMiddleware(), stepController.DeleteStepById)
}
