package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive")


type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	DOB         string             `json:"dob" bson:"dob,omitempty"`
	Address     string             `json:"address" bson:"address,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	CreatedAt   time.Time             `json:"ca" bson:"ca,omitempty"`
	UpdatedAt   time.Time             `json:"ua" bson:"ua,omitempty"`
}
