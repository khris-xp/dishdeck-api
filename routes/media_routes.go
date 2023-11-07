package routes

import (
	"github.com/gofiber/fiber/v2"

	"dishdeck-api/controllers"
	"dishdeck-api/middlewares"
	"dishdeck-api/repositories"
)

func MediaRoutes(app *fiber.App) {

	mediaRepo := repositories.NewMediaRepository()
	mediaController := controllers.NewMediaController(mediaRepo)

	app.Post("/api/file", middlewares.AuthMiddleware(), mediaController.FileUpload)
	app.Post("/api/remote", middlewares.AuthMiddleware(), mediaController.RemoteUpload)
}
