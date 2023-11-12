package models

import "mime/multipart"

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}

type Media struct {
	StatusCode int     `json:"statusCode"`
	Message    string  `json:"message"`
	Data       *string `json:"data"`
}
