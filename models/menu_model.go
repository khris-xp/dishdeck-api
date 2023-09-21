package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Menu struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    Name     string             `json:"name,omitempty" validate:"required"`
    Description string             `json:"description,omitempty" validate:"required"`
	ImageUrl string             `json:"imageUrl,omitempty" validate:"required"`
    Category    string             `json:"category,omitempty" validate:"required"`
}
