package responses

import (
	"github.com/gofiber/fiber/v2"

	"dishdeck-api/types"
)

func GetBlogSuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.BlogListResponse{
		Status:  statusCode,
		Message: "success",
		Data: data,
	})
}

func CreateBlogSuccessResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.BlogResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Created Successfully",
	})
}

func UpdateBlogSuccessResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.BlogResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Updated Successfully",
	})
}

func DeleteBlogSuccessResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.BlogResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Deleted Successfully",
	})
}
