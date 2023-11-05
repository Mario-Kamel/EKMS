package repositories

import (
	"context"
	"fmt"

	"github.com/Mario-Kamel/EKMS/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssignmentRepoInterface interface {
	GetAllAssignments(ctx context.Context) ([]models.Assignment, error)
	GetAssignmentById(ctx context.Context, id string) (*models.Assignment, error)
	CreateAssignment(ctx context.Context, assignment models.Assignment) (*models.Assignment, error)
	UpdateAssignment(ctx context.Context, id string, assignment models.Assignment) (*models.Assignment, error)
	DeleteAssignment(ctx context.Context, id string) error
	AddSubmission(ctx context.Context, assignmentID primitive.ObjectID, submittedID primitive.ObjectID) (*models.Assignment, error)
}

type AssignmentRepo struct {
	db *mongo.Client
}

func NewAssignmentRepo(db *mongo.Client) *AssignmentRepo {
	return &AssignmentRepo{
		db: db,
	}
}

func (m *AssignmentRepo) GetAllAssignments(ctx context.Context) ([]models.Assignment, error) {
	assignments := []models.Assignment{}
	cur, err := m.db.Database("ekms").Collection("assignments").Find(ctx, bson.D{})
	if err != nil {
		fmt.Printf("Error while getting all assignments: %v\n", err)
		return nil, err
	}

	for cur.Next(ctx) {
		var assignment models.Assignment
		err := cur.Decode(&assignment)
		if err != nil {
			fmt.Printf("Error while decoding assignment: %v\n", err)
			return nil, err
		}
		assignments = append(assignments, assignment)
	}

	return assignments, nil
}

func (m *AssignmentRepo) GetAssignmentById(ctx context.Context, id string) (*models.Assignment, error) {
	var assignment models.Assignment
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting id to object id: %v\n", err)
		return nil, err
	}

	err = m.db.Database("ekms").Collection("assignments").FindOne(ctx, bson.M{"_id": oid}).Decode(&assignment)
	if err != nil {
		fmt.Printf("Error while getting assignment by id: %v\n", err)
		return nil, err
	}

	return &assignment, nil
}

func (m *AssignmentRepo) CreateAssignment(ctx context.Context, assignment models.Assignment) (*models.Assignment, error) {
	res, err := m.db.Database("ekms").Collection("assignments").InsertOne(ctx, assignment)
	if err != nil {
		fmt.Printf("Error while creating assignment: %v\n", err)
		return nil, err
	}

	assignment.ID = res.InsertedID.(primitive.ObjectID)

	return &assignment, nil
}

func (m *AssignmentRepo) UpdateAssignment(ctx context.Context, id string, assignment models.Assignment) (*models.Assignment, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting id to object id: %v\n", err)
		return nil, err
	}

	_, err = m.db.Database("ekms").Collection("assignments").UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": assignment})
	if err != nil {
		fmt.Printf("Error while updating assignment: %v\n", err)
		return nil, err
	}

	return &assignment, nil
}

func (m *AssignmentRepo) DeleteAssignment(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting id to object id: %v\n", err)
		return err
	}

	_, err = m.db.Database("ekms").Collection("assignments").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		fmt.Printf("Error while deleting assignment: %v\n", err)
		return err
	}

	return nil
}

// Add submission function
func (m *AssignmentRepo) AddSubmission(ctx context.Context, assignmentID primitive.ObjectID, submittedID primitive.ObjectID) (*models.Assignment, error) {
	var assignment models.Assignment
	_, err := m.db.Database("ekms").Collection("assignments").UpdateOne(ctx, bson.M{"_id": assignmentID}, bson.M{"$push": bson.M{"submittedIds": submittedID}})
	if err != nil {
		fmt.Printf("Error while updating assignment: %v\n", err)
		return nil, err
	}

	err = m.db.Database("ekms").Collection("assignments").FindOne(ctx, bson.M{"_id": assignmentID}).Decode(&assignment)
	if err != nil {
		fmt.Printf("Error while getting assignment by id: %v\n", err)
		return nil, err
	}

	return &assignment, nil
}
