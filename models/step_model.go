package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Step struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Content   string             `json:"content,omitempty" validate:"required"`
	MenuId    primitive.ObjectID `json:"menuId,omitempty" bson:"menuId"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
