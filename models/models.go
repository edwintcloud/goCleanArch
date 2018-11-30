package models

import (
	"time"
)

// User model
type User struct {
	ID        interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string      `json:"email" bson:"email"`
	Password  string      `json:"password" bson:"password"`
	UpdatedAt time.Time   `json:"updated_at" bson:"updated_at"`
	CreatedAt time.Time   `json:"created_at" bson:"created_at"`
}

// ResponseError model
type ResponseError struct {
	Error string `json:"error"`
}
