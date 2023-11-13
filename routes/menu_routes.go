package routes

import (
	"dishdeck-api/controllers"
	"dishdeck-api/middlewares"
	"dishdeck-api/repositories"

	"github.com/gofiber/fiber/v2"
)

const MenuAPIPath = "/api/menu/:id"

func MenuRoutes(app *fiber.App) {
	menuRepo := repositories.NewMenuRepository()
	menuController := controllers.NewMenuController(*menuRepo)

	app.Post("/api/menu", middlewares.AuthMiddleware(), menuController.CreateMenu)
	app.Get("/api/menu", menuController.GetAllMenu)
	app.Get(MenuAPIPath, menuController.GetMenuByID)
	app.Put(MenuAPIPath, middlewares.AuthMiddleware(), menuController.UpdateMenuByID)
	app.Delete(MenuAPIPath, middlewares.AuthMiddleware(), menuController.DeleteMenuByID)
	app.Put(MenuAPIPath+"/like", middlewares.AuthMiddleware(), menuController.LikedMenu)
	app.Put(MenuAPIPath+"/unlike", middlewares.AuthMiddleware(), menuController.UnlikedMenu)
	app.Put(MenuAPIPath+"/rating", middlewares.AuthMiddleware(), menuController.EditRatingMenu)
	app.Put(MenuAPIPath+"/review", middlewares.AuthMiddleware(), menuController.EditReviewMenu)
}
