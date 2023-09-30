package responses

import (
	"github.com/gofiber/fiber/v2"

	"dishdeck-api/types"
)

func GetMenuSuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.MenuListResponse{
		Status:  statusCode,
		Message: "success",
		Data:    &fiber.Map{"data": data},
	})
}

func CreateMenuSuccessResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.MenuResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Created Successfully",
	})
}

func UpdateMenuSuccessResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.MenuResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Updated Successfully",
	})
}

func DeleteMenuSuccessResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.MenuResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Deleted Successfully",
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, errMsg string) error {
	return c.Status(statusCode).JSON(types.MenuListResponse{
		Status:  statusCode,
		Message: "error",
		Data:    &fiber.Map{"data": errMsg},
	})
}
