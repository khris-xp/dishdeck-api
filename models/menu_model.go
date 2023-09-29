package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
	ImageUrl    string             `json:"imageUrl,omitempty" validate:"required"`
	Category    string             `json:"category,omitempty" validate:"required"`
	Review      float64            `json:"review" validate:"required"`
	Rate        float64            `json:"rate" validate:"required"`
	CreatedBy   string             `json:"createdBy,omitempty"`
	Likes       int                `json:"likes,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}
