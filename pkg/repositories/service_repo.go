package repositories

import (
	"context"
	"fmt"

	"github.com/Mario-Kamel/EKMS/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceRepoInterface interface {
	GetAllServices(ctx context.Context) ([]models.Service, error)
	GetServiceById(ctx context.Context, id string) (*models.Service, error)
	CreateService(ctx context.Context, service models.Service) (*models.Service, error)
	UpdateService(ctx context.Context, id string, service models.Service) (*models.Service, error)
	DeleteService(ctx context.Context, id string) error

	AddAttendanceRecord(ctx context.Context, serviceID primitive.ObjectID, ar models.AttendanceRecord) (*models.Service, error)
	EditAttendanceRecord(ctx context.Context, serviceID primitive.ObjectID, ar models.AttendanceRecord) (*models.Service, error)
}

type ServiceRepo struct {
	db *mongo.Client
}

func NewServiceRepo(db *mongo.Client) *ServiceRepo {
	return &ServiceRepo{
		db: db,
	}
}

func (m *ServiceRepo) GetAllServices(ctx context.Context) ([]models.Service, error) {
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

func (m *ServiceRepo) GetServiceById(ctx context.Context, id string) (*models.Service, error) {
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

func (m *ServiceRepo) CreateService(ctx context.Context, service models.Service) (*models.Service, error) {
	service.ID = primitive.NewObjectID()
	_, err := m.db.Database("ekms").Collection("services").InsertOne(ctx, service)
	if err != nil {
		fmt.Printf("Error while creating service: %v\n", err)
		return nil, err
	}

	return &service, nil
}

func (m *ServiceRepo) UpdateService(ctx context.Context, id string, service models.Service) (*models.Service, error) {
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

func (m *ServiceRepo) DeleteService(ctx context.Context, id string) error {
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

func (m *ServiceRepo) AddAttendanceRecord(ctx context.Context, serviceID primitive.ObjectID, ar models.AttendanceRecord) (*models.Service, error) {
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

func (m *ServiceRepo) EditAttendanceRecord(ctx context.Context, serviceID primitive.ObjectID, ar models.AttendanceRecord) (*models.Service, error) {
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
