package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `json:"title" validate:"required"`
	Content   string             `json:"content" validate:"required"`
	Author    string             `json:"author" validate:"required"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}
