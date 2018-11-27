package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// User model
type User struct {
	DocumentID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	ID         int64         `json:"id,omitempty"`
	Email      string        `json:"email" bson:"email"`
	Password   string        `json:"password" bson:"password"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
}
