package repositories

import (
	"context"
	"fmt"

	"github.com/Mario-Kamel/EKMS/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryInterface interface {
	GetAllPersons(ctx context.Context) ([]models.Person, error)
	GetPersonById(ctx context.Context, id string) (*models.Person, error)
	CreatePerson(ctx context.Context, person models.Person) (*models.Person, error)
	UpdatePerson(ctx context.Context, id string, person models.Person) (*models.Person, error)
	DeletePerson(ctx context.Context, id string) error

	GetAllServices(ctx context.Context) ([]models.Service, error)
	GetServiceById(ctx context.Context, id string) (*models.Service, error)
	CreateService(ctx context.Context, service models.Service) (*models.Service, error)
	UpdateService(ctx context.Context, id string, service models.Service) (*models.Service, error)
	DeleteService(ctx context.Context, id string) error

	AddAttendanceRecord(ctx context.Context, serviceID primitive.ObjectID, ar models.AttendanceRecord) (*models.Service, error)
	EditAttendanceRecord(ctx context.Context, serviceID primitive.ObjectID, ar models.AttendanceRecord) (*models.Service, error)
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
	cur, err := m.db.Database("ekms").Collection("people").Find(ctx, bson.D{})
	if err != nil {
		fmt.Printf("Error while getting all persons: %v\n", err)
		return nil, err
	}

	for cur.Next(ctx) {
		var person models.Person
		err := cur.Decode(&person)
		//TODO REmove
		fmt.Println("Person: ", person)
		if err != nil {
			fmt.Printf("Error while decoding person: %v\n", err)
		}
		persons = append(persons, person)
	}

	return persons, nil
}

func (m *MongoRepo) GetPersonById(ctx context.Context, id string) (*models.Person, error) {
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

func (m *MongoRepo) CreatePerson(ctx context.Context, person models.Person) (*models.Person, error) {
	person.ID = primitive.NewObjectID()
	_, err := m.db.Database("ekms").Collection("people").InsertOne(ctx, person)
	if err != nil {
		fmt.Printf("Error while creating person: %v\n", err)
		return nil, err
	}

	return &person, nil
}

func (m *MongoRepo) UpdatePerson(ctx context.Context, id string, person models.Person) (*models.Person, error) {
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

func (m *MongoRepo) DeletePerson(ctx context.Context, id string) error {
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

func (m *MongoRepo) GetAllServices(ctx context.Context) ([]models.Service, error) {
	services := []models.Service{}
	cur, err := m.db.Database("ekms").Collection("services").Find(ctx, bson.D{})
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

func (m *MongoRepo) GetServiceById(ctx context.Context, id string) (*models.Service, error) {
	var service models.Service
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting id to object id: %v\n", err)
		return nil, err
	}
	err = m.db.Database("ekms").Collection("services").FindOne(ctx, bson.M{"_id": oid}).Decode(&service)
	if err != nil {
		fmt.Printf("Error while getting service by id: %v\n", err)
		return nil, err
	}

	return &service, nil
}

func (m *MongoRepo) CreateService(ctx context.Context, service models.Service) (*models.Service, error) {
	service.ID = primitive.NewObjectID()
	_, err := m.db.Database("ekms").Collection("services").InsertOne(ctx, service)
	if err != nil {
		fmt.Printf("Error while creating service: %v\n", err)
		return nil, err
	}

	return &service, nil
}

func (m *MongoRepo) UpdateService(ctx context.Context, id string, service models.Service) (*models.Service, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	service.ID = oid
	if err != nil {
		fmt.Printf("Error while converting id to object id: %v\n", err)
		return nil, err
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "date", Value: service.Date},
			{Key: "subject", Value: service.Subject},
			{Key: "speaker", Value: service.Speaker},
			{Key: "bibleChapter", Value: service.BibleChapter},
		}},
	}

	_, err = m.db.Database("ekms").Collection("services").UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		fmt.Printf("Error while updating service: %v\n", err)
		return nil, err
	}

	return &service, nil
}

func (m *MongoRepo) DeleteService(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting id to object id: %v\n", err)
		return err
	}
	_, err = m.db.Database("ekms").Collection("services").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		fmt.Printf("Error while deleting service: %v\n", err)
		return err
	}

	return nil
}

func (m *MongoRepo) AddAttendanceRecord(ctx context.Context, serviceID primitive.ObjectID, ar models.AttendanceRecord) (*models.Service, error) {
	fmt.Println("Attendance record: ", ar)
	_, err := m.db.Database("ekms").Collection("services").UpdateOne(ctx, bson.M{"_id": serviceID}, bson.M{"$push": bson.M{"attendanceRecord": ar}})
	if err != nil {
		fmt.Printf("Error while adding attendance record: %v\n", err)
		return nil, err
	}

	service, err := m.GetServiceById(ctx, serviceID.Hex())
	if err != nil {
		fmt.Printf("Error while getting service by id: %v\n", err)
		return nil, err
	}

	return service, nil
}

func (m *MongoRepo) EditAttendanceRecord(ctx context.Context, serviceID primitive.ObjectID, ar models.AttendanceRecord) (*models.Service, error) {
	//Replace the attendance record in the service having id = ar.ServiceID and having attendanceRecord.personId = ar.PersonID with ar
	_, err := m.db.Database("ekms").Collection("services").UpdateOne(ctx, bson.M{"_id": serviceID, "attendanceRecord.personId": ar.PersonID}, bson.M{"$set": bson.M{"attendanceRecord.$": ar}})
	if err != nil {
		fmt.Printf("Error while editing attendance record: %v\n", err)
		return nil, err
	}

	service, err := m.GetServiceById(ctx, serviceID.Hex())
	if err != nil {
		fmt.Printf("Error while getting service by id: %v\n", err)
		return nil, err
	}

	return service, nil

}
