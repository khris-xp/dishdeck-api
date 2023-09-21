package routes

import (
	"dishdeck-api/controllers"
	"dishdeck-api/repositories"

	"github.com/gofiber/fiber/v2"
)

func MenuRoutes(app *fiber.App) {
	menuRepo := repositories.NewMenuRepository()
	menuController := controllers.NewMenuController(menuRepo)

	app.Post("/api/menu", menuController.CreateMenu)
	app.Get("/api/menu", menuController.GetAllMenu)
	app.Get("/api/menu/:id", menuController.GetMenuByID)
}
