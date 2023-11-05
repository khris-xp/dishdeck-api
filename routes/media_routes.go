package routes

import (
	"github.com/gofiber/fiber/v2"

	"dishdeck-api/controllers"
	"dishdeck-api/middlewares"
)

func MediaRoutes(app *fiber.App) {

	app.Post("/api/file", middlewares.AuthMiddleware(), controllers.FileUpload)
	app.Post("/api/remote", middlewares.AuthMiddleware(), controllers.RemoteUpload)
}
