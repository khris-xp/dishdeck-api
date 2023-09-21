package responses

import (
	"github.com/gofiber/fiber/v2"
)

type MenuResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

func SuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(MenuResponse{
		Status:  statusCode,
		Message: "success",
		Data:    &fiber.Map{"data": data},
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, errMsg string) error {
	return c.Status(statusCode).JSON(MenuResponse{
		Status:  statusCode,
		Message: "error",
		Data:    &fiber.Map{"data": errMsg},
	})
}
