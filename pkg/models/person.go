package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Birthday time.Time          `json:"birthday" bson:"birthday,omitempty"`
	Phone    string             `json:"phone" bson:"phone,omitempty"`
	Address  string             `json:"address" bson:"address,omitempty"`
	Fr       string             `json:"fr" bson:"fr,omitempty"`
	Degree   string             `json:"degree" bson:"degree,omitempty"`
}
