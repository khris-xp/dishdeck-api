package controllers

import (
	"dishdeck-api/models"
	"dishdeck-api/repositories"
	"dishdeck-api/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MediaController struct {
	MediaRepo *repositories.MediaRepository
}

func NewMediaController(mediaRepo *repositories.MediaRepository) *MediaController {
	return &MediaController{MediaRepo: mediaRepo}
}

func (m *MediaController) FileUpload(c *fiber.Ctx) error {
	formHeader, err := c.FormFile("file")
	if err != nil {
		return responses.UploadMediaErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return responses.UploadMediaErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	uploadUrl, err := m.MediaRepo.FileUpload(models.File{File: formFile})
	if err != nil {
		return responses.UploadMediaErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return responses.UploadMediaSuccessResponse(c, http.StatusOK, uploadUrl)
}

func (m *MediaController) RemoteUpload(c *fiber.Ctx) error {
	var url models.Url

	if err := c.BodyParser(&url); err != nil {
		return responses.UploadMediaErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	uploadUrl, err := m.MediaRepo.RemoteUpload(url)
	if err != nil {
		return responses.UploadMediaErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return responses.UploadMediaSuccessResponse(c, http.StatusOK, uploadUrl)
}
