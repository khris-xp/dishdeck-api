package routes

import (
	"dishdeck-api/controllers"
	"dishdeck-api/middlewares"
	"dishdeck-api/repositories"

	"github.com/gofiber/fiber/v2"
)

func MenuRoutes(app *fiber.App) {
	menuRepo := repositories.NewMenuRepository()
	menuController := controllers.NewMenuController(menuRepo)

	app.Post("/api/menu", middlewares.AuthMiddleware(), menuController.CreateMenu)
	app.Get("/api/menu", menuController.GetAllMenu)
	app.Get("/api/menu/:id", menuController.GetMenuByID)
	app.Put("/api/menu/:id", middlewares.AuthMiddleware(), menuController.UpdateMenuByID)
	app.Delete("/api/menu/:id", middlewares.AuthMiddleware(), menuController.DeleteMenuByID)
	app.Put("/api/menu/:id/like", middlewares.AuthMiddleware(), menuController.LikedMenu)
	app.Put("/api/menu/:id/unlike", middlewares.AuthMiddleware(), menuController.UnlikedMenu)
	app.Put("/api/menu/:id/rating", middlewares.AuthMiddleware(), menuController.EditRatingMenu)
	app.Put("/api/menu/:id/review", middlewares.AuthMiddleware(), menuController.EditReviewMenu)
}
