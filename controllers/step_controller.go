package controllers

import (
	"dishdeck-api/models"
	"dishdeck-api/repositories"
	"dishdeck-api/responses"
	"dishdeck-api/types"

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

func (sc *StepController) GetStepById(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	step, err := sc.StepRepo.GetStepById(c.Context(), id)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(c, fiber.StatusOK, step)
}

func (sc *StepController) UpdateStep(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var reqBody types.StepUpdateRequest

	if err := c.BodyParser(&reqBody); err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if validationErr := validate.Struct(&reqBody); validationErr != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
	}

	existingStep, err := sc.StepRepo.GetStepById(c.Context(), id)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	existingStep.Title = reqBody.Title
	existingStep.Content = reqBody.Content

	err = sc.StepRepo.UpdateStepById(c.Context(), id, existingStep)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return responses.SuccessResponse(c, fiber.StatusOK, nil)
}

func (sc *StepController) DeleteStepById(ctx *fiber.Ctx) error {
	stepId := ctx.Params("id")

	id, err := primitive.ObjectIDFromHex(stepId)

	if err != nil {
		return responses.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	err = sc.StepRepo.DeleteStepById(ctx.Context(), id)

	if err != nil {
		return responses.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return responses.SuccessResponse(ctx, fiber.StatusOK, nil)
}
