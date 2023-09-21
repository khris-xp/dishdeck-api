package routes

import (
	"dishdeck-api/controllers"
	"dishdeck-api/repositories"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userRepo := repositories.NewUserRepository()
	userController := controllers.NewAuthController(userRepo)

	app.Post("/register", userController.Register)
	app.Post("/login", userController.Login)
	app.Get("/user", userController.GetUserProfile)
}
