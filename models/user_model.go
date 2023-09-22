package models

import "time"

type User struct {
	ID          string    `bson:"_id,omitempty"`
	Username    string    `json:"username" bson:"username"`
	Email       string    `json:"email" bson:"email"`
	Password    string    `json:"password" bson:"password"`
	UserProfile string    `json:"userProfile" bson:"userProfile"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
