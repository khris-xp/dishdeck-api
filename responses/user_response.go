package responses

import (
	"github.com/gofiber/fiber/v2"

	"dishdeck-api/types"
)

func AddWishListResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.WishListResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Added Successfully",
	})
}

func RemoveWishListResponse(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(types.WishListResponse{
		Status:  statusCode,
		Message: "success",
		Data:    "Removed Successfully",
	})
}
