package responses

import (
	"github.com/gofiber/fiber/v2"

	"dishdeck-api/types"
)

func UploadMediaSuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(types.MediaListResponse{
		Status:  statusCode,
		Message: "success",
		Data:    data,
	})
}

func UploadMediaErrorResponse(c *fiber.Ctx, statusCode int, errMsg string) error {
	return c.Status(statusCode).JSON(types.MediaListResponse{
		Status:  statusCode,
		Message: "error",
		Data:    errMsg,
	})
}
