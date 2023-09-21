// controllers/menu_controller.go
package controllers

import (
	"dishdeck-api/models"
	"dishdeck-api/repositories"
	"dishdeck-api/responses"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	validate = validator.New()
)

type MenuController struct {
	MenuRepo *repositories.MenuRepository
}

func NewMenuController(menuRepo *repositories.MenuRepository) *MenuController {
	return &MenuController{MenuRepo: menuRepo}
}

func (mc *MenuController) CreateMenu(c *fiber.Ctx) error {
	var menu models.Menu
	if err := c.BodyParser(&menu); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if validationErr := validate.Struct(&menu); validationErr != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, validationErr.Error())
	}

	menuID, err := mc.MenuRepo.CreateMenu(c.Context(), menu)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, http.StatusCreated, menuID)
}

func (mc *MenuController) GetAllMenu(c *fiber.Ctx) error {
	menuList, err := mc.MenuRepo.GetAllMenu(c.Context())
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, http.StatusOK, menuList)
}

func (mc *MenuController) GetMenuByID(c *fiber.Ctx) error {
	menuID := c.Params("id")
	id, err := primitive.ObjectIDFromHex(menuID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	menu, err := mc.MenuRepo.GetMenuByID(c.Context(), id)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, http.StatusOK, menu)
}
