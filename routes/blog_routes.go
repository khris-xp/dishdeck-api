package routes

import (
	"github.com/gofiber/fiber/v2"

	"dishdeck-api/controllers"
	"dishdeck-api/middlewares"
	"dishdeck-api/repositories"
)

func BlogRoutes(app *fiber.App) {
	blogRepo := repositories.NewBlogRepository()
	blogController := controllers.NewBlogController(blogRepo)

	app.Post("/api/blog", middlewares.AuthMiddleware(), blogController.CreateBlog)
	app.Get("/api/blog", blogController.GetAllBlog)
	app.Get("/api/blog/:id", blogController.GetBlogById)
	app.Put("/api/blog/:id", middlewares.AuthMiddleware(), blogController.UpdateBlogById)
	app.Delete("/api/blog/:id", middlewares.AuthMiddleware(), blogController.DeleteBlogById)
}
