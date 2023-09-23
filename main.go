package main

import (
	"github.com/gofiber/fiber/v2"

	"dishdeck-api/configs"
	"dishdeck-api/routes"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello, from Disdeck API"})
	})

	configs.ConnectDB()
	routes.MenuRoutes(app)
	routes.UserRoutes(app)
	routes.StepRoutes(app)
	port := configs.EnvPort()
	app.Listen(":" + port)
}
