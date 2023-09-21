package controllers

import (
	"context"
	"dishdeck-api/configs"
	"dishdeck-api/models"
	"dishdeck-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = configs.GetCollection(configs.DB, "menu")
var validate = validator.New()

func CreateMenu(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var menu models.Menu
	defer cancel()

	if err := c.BodyParser(&menu); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MenuResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&menu); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MenuResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newMenu := models.Menu{
		Id:          primitive.NewObjectID(),
		Name:        menu.Name,
		Description: menu.Description,
		ImageUrl:    menu.ImageUrl,
		Category:    menu.Category,
	}

	result, err := menuCollection.InsertOne(ctx, newMenu)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MenuResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.MenuResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllMenu(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var menu []models.Menu
	defer cancel()

	results, err := menuCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MenuResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleMenu models.Menu
		if err = results.Decode(&singleMenu); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.MenuResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		menu = append(menu, singleMenu)
	}

	return c.Status(http.StatusOK).JSON(
		responses.MenuResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": menu}},
	)

}
