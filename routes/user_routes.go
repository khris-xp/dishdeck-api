package routes

import (
	"dishdeck-api/controllers"
	"dishdeck-api/middlewares"
	"dishdeck-api/repositories"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userRepo := repositories.NewUserRepository()
	userController := controllers.NewAuthController(userRepo)

	app.Post("/auth/register", userController.Register)
	app.Post("auth/login", userController.Login)
	app.Get("/auth/user", userController.GetUserProfile)
	app.Get("/auth/user/:id", userController.GetUserById)
	app.Put("/auth/user/:id/wish/:menuId", middlewares.AuthMiddleware(), userController.AddWishListByMenuID)
	app.Put("/auth/user/:id/remove-wish/:menuId", middlewares.AuthMiddleware(), userController.RemoveWishListByMenuID)
}
