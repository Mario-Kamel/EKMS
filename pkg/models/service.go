package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ID           primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Date         time.Time            `json:"date" bson:"date,omitempty"`
	Subject      string               `json:"subject" bson:"subject,omitempty"`
	Speaker      string               `json:"speaker" bson:"speaker,omitempty"`
	BibleChapter string               `json:"bibleChapter" bson:"bibleChapter,omitempty"`
	Attendants   []primitive.ObjectID `json:"attendants" bson:"attendants,omitempty"`
}
