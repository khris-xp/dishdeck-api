package controllers

import (
	"dishdeck-api/models"
	"dishdeck-api/repositories"
	"dishdeck-api/responses"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StepController struct {
	StepRepo *repositories.StepRepository
}

func NewStepController(stepRepo *repositories.StepRepository) *StepController {
	return &StepController{StepRepo: stepRepo}
}

func (sc *StepController) CreateStep(c *fiber.Ctx) error {
	var step models.Step
	menuId := c.Params("id")

	id, err := primitive.ObjectIDFromHex(menuId)

	if err := c.BodyParser(&step); err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	stepID, err := sc.StepRepo.CreateStep(c.Context(), step, id)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, fiber.StatusCreated, stepID)
}

func (sc *StepController) GetAllStep(c *fiber.Ctx) error {
	step, err := sc.StepRepo.GetAllStep(c.Context())
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, fiber.StatusOK, step)
}

func (sc *StepController) GetStepByMenuId(c *fiber.Ctx) error {
	menuIdStr := c.Params("id")

	menuId, err := primitive.ObjectIDFromHex(menuIdStr)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	steps, err := sc.StepRepo.GetStepByMenuId(c.Context(), menuId)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, fiber.StatusOK, steps)
}
