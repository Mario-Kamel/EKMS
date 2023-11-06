package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Assignment struct {
	ID          primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	ServiceID   primitive.ObjectID     `json:"serviceId" bson:"serviceId,omitempty"`
	Title       string                 `json:"title" bson:"title,omitempty"`
	Deadline    time.Time              `json:"deadline" bson:"deadline,omitempty"`
	Submissions []AssignmentSubmission `json:"submissions" bson:"submissions,omitempty"`
}

type AssignmentSubmission struct {
	PersonID primitive.ObjectID `json:"personId" bson:"personId,omitempty"`
	Time     time.Time          `json:"time" bson:"time,omitempty"`
}
