package controllers

import (
	"dishdeck-api/models"
	"dishdeck-api/repositories"
	"dishdeck-api/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func FileUpload(c *fiber.Ctx) error {
	formHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       nil,
			})
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       nil,
			})
	}

	uploadUrl, err := repositories.NewMediaUpload().FileUpload(models.File{File: formFile})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       nil,
			})
	}

	return c.Status(http.StatusOK).JSON(
		responses.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &uploadUrl,
		})
}

func RemoteUpload(c *fiber.Ctx) error {
	var url models.Url

	if err := c.BodyParser(&url); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.MediaDto{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Data:       nil,
			})
	}

	uploadUrl, err := repositories.NewMediaUpload().RemoteUpload(url)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       nil,
			})
	}

	return c.Status(http.StatusOK).JSON(
		responses.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &uploadUrl,
		})
}
