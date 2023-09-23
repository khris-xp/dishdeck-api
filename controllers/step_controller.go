package controllers

import (
	"dishdeck-api/models"
	"dishdeck-api/repositories"
	"dishdeck-api/responses"

	"github.com/gofiber/fiber/v2"
)

type StepController struct {
	StepRepo *repositories.StepRepository
}

func NewStepController(stepRepo *repositories.StepRepository) *StepController {
	return &StepController{StepRepo: stepRepo}
}

func (sc *StepController) CreateStep(c *fiber.Ctx) error {
	var step models.Step
	var menu models.Menu

	if err := c.BodyParser(&step); err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	stepID, err := sc.StepRepo.CreateStep(c.Context(), step, menu)
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
