package types

import (
	"github.com/gofiber/fiber/v2"
)

type MenuResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

type MenuUpdateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"imageUrl" validate:"required"`
	Category    string `json:"category" validate:"required"`
}