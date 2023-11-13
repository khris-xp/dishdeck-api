package repositories

import (
	"dishdeck-api/helper"
	"dishdeck-api/models"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type MediaRepositoryInterface interface {
	FileUpload(file models.File) (string, error)
	RemoteUpload(url models.Url) (string, error)
}

type MediaRepository struct{}

func NewMediaRepository() *MediaRepository {
	return &MediaRepository{}
}

func (*MediaRepository) FileUpload(file models.File) (string, error) {
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	uploadUrl, err := helper.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*MediaRepository) RemoteUpload(url models.Url) (string, error) {
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	uploadUrl, errUrl := helper.ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}
