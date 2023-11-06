package service

import (
	"context"

	"github.com/Mario-Kamel/EKMS/pkg/models"
	"github.com/Mario-Kamel/EKMS/pkg/repositories"
)

type AssignmentService struct {
	repo repositories.AssignmentRepoInterface
}

func NewAssignmentService(repo repositories.AssignmentRepoInterface) *AssignmentService {
	return &AssignmentService{
		repo: repo,
	}
}

func (s *AssignmentService) GetAllAssignments(ctx context.Context) ([]models.Assignment, error) {
	assignments, err := s.repo.GetAllAssignments(ctx)
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

func (s *AssignmentService) GetAssignmentById(ctx context.Context, id string) (*models.Assignment, error) {
	assignment, err := s.repo.GetAssignmentById(ctx, id)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

func (s *AssignmentService) CreateAssignment(ctx context.Context, assignment models.Assignment) (*models.Assignment, error) {
	a, err := s.repo.CreateAssignment(ctx, assignment)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AssignmentService) UpdateAssignment(ctx context.Context, id string, assignment models.Assignment) (*models.Assignment, error) {
	a, err := s.repo.UpdateAssignment(ctx, id, assignment)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AssignmentService) DeleteAssignment(ctx context.Context, id string) error {
	err := s.repo.DeleteAssignment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
