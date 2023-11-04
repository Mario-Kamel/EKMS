package service

import (
	"context"

	"github.com/Mario-Kamel/EKMS/pkg/models"
	"github.com/Mario-Kamel/EKMS/pkg/repositories"
)

type ServiceService struct {
	repo repositories.RepositoryInterface
}

func NewServiceService(repo repositories.RepositoryInterface) *ServiceService {
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
