package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Date             time.Time          `json:"date" bson:"date,omitempty"`
	Subject          string             `json:"subject" bson:"subject,omitempty"`
	Speaker          string             `json:"speaker" bson:"speaker,omitempty"`
	BibleChapter     string             `json:"bibleChapter" bson:"bibleChapter,omitempty"`
	AttendanceRecord []AttendanceRecord `json:"attendanceRecord" bson:"attendanceRecord,omitempty"`
}

type AttendanceRecord struct {
	ServiceID primitive.ObjectID `json:"serviceId" bson:"serviceId,omitempty"`
	PersonID  primitive.ObjectID `json:"personId" bson:"personId,omitempty"`
	Time      time.Time          `json:"time" bson:"time,omitempty"`
	Status    string             `json:"status" bson:"status,omitempty"`
}
