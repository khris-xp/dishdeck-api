package responses

import (
	"github.com/gofiber/fiber/v2"

	"dishdeck-api/types"
)

func SuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.MenuResponse{
		Status:  statusCode,
		Message: "success",
		Data:    &fiber.Map{"data": data},
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, errMsg string) error {
	return c.Status(statusCode).JSON(types.MenuResponse{
		Status:  statusCode,
		Message: "error",
		Data:    &fiber.Map{"data": errMsg},
	})
}
