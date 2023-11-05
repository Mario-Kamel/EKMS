package service

import (
	"context"
	"fmt"

	"github.com/Mario-Kamel/EKMS/pkg/models"
	"github.com/Mario-Kamel/EKMS/pkg/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceService struct {
	repo repositories.ServiceRepoInterface
}

func NewServiceService(repo repositories.ServiceRepoInterface) *ServiceService {
	return &ServiceService{
		repo: repo,
	}
}

func (s *ServiceService) GetAllServices(ctx context.Context) ([]models.Service, error) {
	services, err := s.repo.GetAllServices(ctx)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (s *ServiceService) GetServiceById(ctx context.Context, id string) (*models.Service, error) {
	service, err := s.repo.GetServiceById(ctx, id)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (s *ServiceService) CreateService(ctx context.Context, service models.Service) (*models.Service, error) {
	serv, err := s.repo.CreateService(ctx, service)
	if err != nil {
		return nil, err
	}
	return serv, nil
}

func (s *ServiceService) UpdateService(ctx context.Context, id string, service models.Service) (*models.Service, error) {
	serv, err := s.repo.UpdateService(ctx, id, service)
	if err != nil {
		return nil, err
	}
	return serv, nil
}

func (s *ServiceService) DeleteService(ctx context.Context, id string) error {
	err := s.repo.DeleteService(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceService) AddAttendanceRecord(ctx context.Context, serviceID string, ar models.AttendanceRecord) (*models.Service, error) {
	oid, err := primitive.ObjectIDFromHex(serviceID)
	if err != nil {
		fmt.Println("Error while converting id to object id: ", err)
		return nil, err
	}
	serv, err := s.repo.AddAttendanceRecord(ctx, oid, ar)
	if err != nil {
		return nil, err
	}
	return serv, nil
}

func (s *ServiceService) EditAttendanceRecord(ctx context.Context, serviceID string, ar models.AttendanceRecord) (*models.Service, error) {
	oid, err := primitive.ObjectIDFromHex(serviceID)
	if err != nil {
		fmt.Println("Error while converting id to object id: ", err)
		return nil, err
	}
	serv, err := s.repo.EditAttendanceRecord(ctx, oid, ar)
	if err != nil {
		return nil, err
	}
	return serv, nil
}

func (s *ServiceService) DeleteAttendanceRecord(ctx context.Context, serviceID string, ar models.AttendanceRecord) (*models.Service, error) {
	oid, err := primitive.ObjectIDFromHex(serviceID)
	if err != nil {
		fmt.Println("Error while converting id to object id: ", err)
		return nil, err
	}
	serv, err := s.repo.DeleteAttendanceRecord(ctx, oid, ar)
	if err != nil {
		return nil, err
	}
	return serv, nil
}
