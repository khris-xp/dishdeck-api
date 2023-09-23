package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"dishdeck-api/configs"
	"dishdeck-api/routes"
)

func loadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
		}
	}
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello, from Disdeck API"})
	})

	configs.ConnectDB()
	loadEnv()
	routes.MenuRoutes(app)
	routes.UserRoutes(app)
	routes.StepRoutes(app)

	port := configs.EnvPort()
	app.Listen(":" + port)
}
