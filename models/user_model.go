package models

import "time"

type User struct {
	Id          string    `bson:"_id,omitempty"`
	Username    string    `json:"username" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	UserProfile string    `json:"userProfile" validate:"required"`
	Wishlist    []string  `json:"wishlist"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
