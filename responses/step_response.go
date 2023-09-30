package responses

import (
	"github.com/gofiber/fiber/v2"

	"dishdeck-api/models"
	"dishdeck-api/types"
)

func GetStepListResponse(c *fiber.Ctx, statusCode int, data []models.Step) error {
	return c.Status(statusCode).JSON(types.StepListResponse{
		Status:  statusCode,
		Message: "success",
		Data:    &fiber.Map{"steps": data},
	})
}

func GetStepResponse(c *fiber.Ctx, statusCode int, data models.Step) error {
	return c.Status(statusCode).JSON(types.StepListResponse{
		Status:  statusCode,
		Message: "success",
		Data:    &fiber.Map{"step": data},
	})
}

func CreateStepSuccessResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.StepResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Created Successfully",
	})
}

func UpdateStepSuccessResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.StepResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Updated Successfully",
	})
}

func DeleteStepSuccessResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.StepResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Deleted Successfully",
	})
}
