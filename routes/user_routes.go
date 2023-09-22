package routes

import (
	"dishdeck-api/controllers"
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
}
