package types

import (
	"github.com/gofiber/fiber/v2"
)

type StepResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

type StepUpdateRequest struct {
	Title        string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
