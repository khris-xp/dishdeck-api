package routes

import (
	"dishdeck-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func MenuRoutes(app *fiber.App) {
	app.Post("/menu", controllers.CreateMenu)
	app.Get("/menu", controllers.GetAllMenu)
}
