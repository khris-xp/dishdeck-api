package models

import "time"

type User struct {
	Id          string    `bson:"_id,omitempty"`
	Username    string    `json:"username" bson:"username" validate:"required"`
	Email       string    `json:"email" bson:"email" validate:"required"`
	Password    string    `json:"password" bson:"password" validate:"required"`
	UserProfile string    `json:"userProfile" bson:"userProfile" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
