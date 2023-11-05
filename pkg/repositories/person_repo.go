package repositories

import (
	"context"
	"fmt"

	"github.com/Mario-Kamel/EKMS/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonRepoInterface interface {
	GetAllPersons(ctx context.Context) ([]models.Person, error)
	GetPersonById(ctx context.Context, id string) (*models.Person, error)
	CreatePerson(ctx context.Context, person models.Person) (*models.Person, error)
	UpdatePerson(ctx context.Context, id string, person models.Person) (*models.Person, error)
	DeletePerson(ctx context.Context, id string) error
}

type PersonRepo struct {
	db *mongo.Client
}

func NewPersonRepo(db *mongo.Client) *PersonRepo {
	return &PersonRepo{
		db: db,
	}
}

func (m *PersonRepo) GetAllPersons(ctx context.Context) ([]models.Person, error) {
	persons := []models.Person{}
	cur, err := m.db.Database("ekms").Collection("people").Find(ctx, bson.D{})
	if err != nil {
		fmt.Printf("Error while getting all persons: %v\n", err)
		return nil, err
	}

	for cur.Next(ctx) {
		var person models.Person
		err := cur.Decode(&person)

		if err != nil {
			fmt.Printf("Error while decoding person: %v\n", err)
		}
		persons = append(persons, person)
	}

	return persons, nil
}

func (m *PersonRepo) GetPersonById(ctx context.Context, id string) (*models.Person, error) {
	var person models.Person
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting id to object id: %v\n", err)
		return nil, err
	}

	err = m.db.Database("ekms").Collection("people").FindOne(ctx, bson.M{"_id": oid}).Decode(&person)
	if err != nil {
		fmt.Printf("Error while getting person by id: %v\n", err)
		return nil, err
	}

	return &person, nil
}

func (m *PersonRepo) CreatePerson(ctx context.Context, person models.Person) (*models.Person, error) {
	person.ID = primitive.NewObjectID()
	_, err := m.db.Database("ekms").Collection("people").InsertOne(ctx, person)
	if err != nil {
		fmt.Printf("Error while creating person: %v\n", err)
		return nil, err
	}

	return &person, nil
}

func (m *PersonRepo) UpdatePerson(ctx context.Context, id string, person models.Person) (*models.Person, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	person.ID = oid
	if err != nil {
		fmt.Printf("Error while converting id to object id: %v\n", err)
		return nil, err
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: person.Name},
			{Key: "birthday", Value: person.Birthday},
			{Key: "phone", Value: person.Phone},
			{Key: "address", Value: person.Address},
			{Key: "fr", Value: person.Fr},
			{Key: "degree", Value: person.Degree},
		}},
	}
	_, err = m.db.Database("ekms").Collection("people").UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		fmt.Printf("Error while updating person: %v\n", err)
		return nil, err
	}
	fmt.Println("Person inside update: ", person)
	return &person, nil
}

func (m *PersonRepo) DeletePerson(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting id to object id: %v\n", err)
		return err
	}
	_, err = m.db.Database("ekms").Collection("people").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		fmt.Printf("Error while deleting person: %v\n", err)
		return err
	}

	return nil
}
