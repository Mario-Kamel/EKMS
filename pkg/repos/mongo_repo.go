package repositories

import (
	"context"
	"fmt"

	"github.com/Mario-Kamel/EKMS/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryInterface interface {
	// Person
	GetAllPersons(ctx context.Context) ([]models.Person, error)
	GetPersonById(ctx context.Context, id string) (*models.Person, error)
	CreatePerson(ctx context.Context, person models.Person) (*models.Person, error)
	UpdatePerson(ctx context.Context, id string, person models.Person) (*models.Person, error)
	DeletePerson(ctx context.Context, id string) error
	// Service
	GetAllServices(ctx context.Context) ([]models.Service, error)
	GetServiceById(ctx context.Context, id string) (*models.Service, error)
	CreateService(ctx context.Context, service models.Service) (*models.Service, error)
	UpdateService(ctx context.Context, id string, service models.Service) (*models.Service, error)
	DeleteService(ctx context.Context, id string) error
}

type MongoRepo struct {
	db *mongo.Client
}

func NewMongoRepo(db *mongo.Client) *MongoRepo {
	return &MongoRepo{
		db: db,
	}
}

func (m *MongoRepo) GetAllPersons(ctx context.Context) ([]models.Person, error) {
	persons := []models.Person{}
	cur, err := m.db.Database("ekms").Collection("people").Find(ctx, nil)
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

func (m *MongoRepo) GetPersonById(ctx context.Context, oid string) (*models.Person, error) {
	var person models.Person
	err := m.db.Database("ekms").Collection("people").FindOne(ctx, bson.M{"_id": oid}).Decode(&person)
	if err != nil {
		fmt.Printf("Error while getting person by id: %v\n", err)
		return nil, err
	}

	return &person, nil
}

func (m *MongoRepo) CreatePerson(ctx context.Context, person models.Person) (*models.Person, error) {
	_, err := m.db.Database("ekms").Collection("people").InsertOne(ctx, person)
	if err != nil {
		fmt.Printf("Error while creating person: %v\n", err)
		return nil, err
	}

	return &person, nil
}

func (m *MongoRepo) UpdatePerson(ctx context.Context, oid string, person models.Person) (*models.Person, error) {
	_, err := m.db.Database("ekms").Collection("people").UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": person})
	if err != nil {
		fmt.Printf("Error while updating person: %v\n", err)
		return nil, err
	}

	return &person, nil
}

func (m *MongoRepo) DeletePerson(ctx context.Context, oid string) error {
	_, err := m.db.Database("ekms").Collection("people").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		fmt.Printf("Error while deleting person: %v\n", err)
		return err
	}

	return nil
}

func (m *MongoRepo) GetAllServices(ctx context.Context) ([]models.Service, error) {
	services := []models.Service{}
	cur, err := m.db.Database("ekms").Collection("services").Find(ctx, nil)
	if err != nil {
		fmt.Printf("Error while getting all services: %v\n", err)
		return nil, err
	}

	for cur.Next(ctx) {
		var service models.Service
		err := cur.Decode(&service)
		if err != nil {
			fmt.Printf("Error while decoding service: %v\n", err)
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}

func (m *MongoRepo) GetServiceById(ctx context.Context, oid string) (*models.Service, error) {
	var service models.Service
	err := m.db.Database("ekms").Collection("services").FindOne(ctx, bson.M{"_id": oid}).Decode(&service)
	if err != nil {
		fmt.Printf("Error while getting service by id: %v\n", err)
		return nil, err
	}

	return &service, nil
}

func (m *MongoRepo) CreateService(ctx context.Context, service models.Service) (*models.Service, error) {
	_, err := m.db.Database("ekms").Collection("services").InsertOne(ctx, service)
	if err != nil {
		fmt.Printf("Error while creating service: %v\n", err)
		return nil, err
	}

	return &service, nil
}

func (m *MongoRepo) UpdateService(ctx context.Context, oid string, service models.Service) (*models.Service, error) {
	_, err := m.db.Database("ekms").Collection("services").UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": service})
	if err != nil {
		fmt.Printf("Error while updating service: %v\n", err)
		return nil, err
	}

	return &service, nil
}

func (m *MongoRepo) DeleteService(ctx context.Context, oid string) error {
	_, err := m.db.Database("ekms").Collection("services").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		fmt.Printf("Error while deleting service: %v\n", err)
		return err
	}

	return nil
}
