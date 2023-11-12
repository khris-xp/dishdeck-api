package routes

import (
    "github.com/gofiber/fiber/v2"
    "dishdeck-api/controllers"
    "dishdeck-api/middlewares"
    "dishdeck-api/repositories"
)

const BlogAPIPath = "/api/blog/:id"

func BlogRoutes(app *fiber.App) {
    blogRepo := repositories.NewBlogRepository()
    blogController := controllers.NewBlogController(blogRepo)

    app.Post("/api/blog", middlewares.AuthMiddleware(), blogController.CreateBlog)
    app.Get("/api/blog", blogController.GetAllBlog)
    app.Get(BlogAPIPath, blogController.GetBlogById)
    app.Put(BlogAPIPath, middlewares.AuthMiddleware(), blogController.UpdateBlogById)
    app.Delete(BlogAPIPath, middlewares.AuthMiddleware(), blogController.DeleteBlogById)
}