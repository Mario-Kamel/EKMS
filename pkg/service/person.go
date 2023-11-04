package service

import (
	"context"

	"github.com/Mario-Kamel/EKMS/pkg/models"
	"github.com/Mario-Kamel/EKMS/pkg/repositories"
)

type PersonService struct {
	repo repositories.RepositoryInterface
}

func NewPersonService(repo repositories.RepositoryInterface) *PersonService {
	return &PersonService{
		repo: repo,
	}
}

func (s *PersonService) GetAllPersons(ctx context.Context) ([]models.Person, error) {
	persons, err := s.repo.GetAllPersons(ctx)
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func (s *PersonService) GetPersonById(ctx context.Context, id string) (*models.Person, error) {
	person, err := s.repo.GetPersonById(ctx, id)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (s *PersonService) CreatePerson(ctx context.Context, person models.Person) (*models.Person, error) {
	p, err := s.repo.CreatePerson(ctx, person)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) UpdatePerson(ctx context.Context, id string, person models.Person) (*models.Person, error) {
	p, err := s.repo.UpdatePerson(ctx, id, person)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) DeletePerson(ctx context.Context, id string) error {
	err := s.repo.DeletePerson(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
