package controllers

import (
	"context"
	"dishdeck-api/configs"
	"dishdeck-api/models"
	"dishdeck-api/repositories"
	"dishdeck-api/responses"
	"dishdeck-api/types"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var (
	validate = validator.New()
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

type MenuController struct {
	MenuRepo *repositories.MenuRepository
}

func NewMenuController(menuRepo *repositories.MenuRepository) *MenuController {
	return &MenuController{MenuRepo: menuRepo}
}

func (mc *MenuController) CreateMenu(c *fiber.Ctx) error {
	var menu models.Menu

	email, exists := c.Locals("email").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if !exists {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "email not found")
	}

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err := c.BodyParser(&menu); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if validationErr := validate.Struct(&menu); validationErr != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, validationErr.Error())
	}

	menuID, err := mc.MenuRepo.CreateMenu(c.Context(), menu, user)
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

func (mc *MenuController) UpdateMenuByID(ctx *fiber.Ctx) error {
	menuID := ctx.Params("id")
	id, err := primitive.ObjectIDFromHex(menuID)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	var reqBody types.MenuUpdateRequest
	if err := ctx.BodyParser(&reqBody); err != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	if validationErr := validate.Struct(&reqBody); validationErr != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, validationErr.Error())
	}

	existingMenu, err := mc.MenuRepo.GetMenuByID(ctx.Context(), id)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	existingMenu.Name = reqBody.Name
	existingMenu.Description = reqBody.Description
	existingMenu.ImageUrl = reqBody.ImageUrl
	existingMenu.Category = reqBody.Category

	err = mc.MenuRepo.UpdateMenuByID(ctx.Context(), id, existingMenu)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(ctx, http.StatusOK, existingMenu)
}

func (mc *MenuController) DeleteMenuByID(ctx *fiber.Ctx) error {
	menuId := ctx.Params("id")

	id, err := primitive.ObjectIDFromHex(menuId)

	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	err = mc.MenuRepo.DeleteMenuByID(ctx.Context(), id)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(ctx, http.StatusOK, nil)
}

func (mc *MenuController) LikedMenu(ctx *fiber.Ctx) error {
	menuId := ctx.Params("id")

	id, err := primitive.ObjectIDFromHex(menuId)

	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	err = mc.MenuRepo.LikedMenu(ctx.Context(), id)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(ctx, http.StatusOK, nil)
}

func (mc *MenuController) UnlikedMenu(ctx *fiber.Ctx) error {
	menuId := ctx.Params("id")

	id, err := primitive.ObjectIDFromHex(menuId)

	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	err = mc.MenuRepo.UnlikedMenu(ctx.Context(), id)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(ctx, http.StatusOK, nil)
}

func (c *MenuController) EditRatingMenu(ctx *fiber.Ctx) error {
	menuId := ctx.Params("id")

	id, err := primitive.ObjectIDFromHex(menuId)

	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	var reqBody types.MenuRatingRequest

	if err := ctx.BodyParser(&reqBody); err != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	if validationErr := validate.Struct(&reqBody); validationErr != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, validationErr.Error())
	}

	err = c.MenuRepo.EditRatingMenu(ctx.Context(), id, reqBody.Rate)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(ctx, http.StatusOK, nil)
}

func (c *MenuController) EditReviewMenu(ctx *fiber.Ctx) error {
	menuId := ctx.Params("id")

	id, err := primitive.ObjectIDFromHex(menuId)

	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	var reqBody types.MenuReviewRequest

	if err := ctx.BodyParser(&reqBody); err != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	if validationErr := validate.Struct(&reqBody); validationErr != nil {
		return responses.ErrorResponse(ctx, http.StatusBadRequest, validationErr.Error())
	}

	err = c.MenuRepo.EditReviewMenu(ctx.Context(), id, reqBody.Review)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(ctx, http.StatusOK, nil)
}