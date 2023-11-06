package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Mario-Kamel/EKMS/pkg/cerrors"
	"github.com/Mario-Kamel/EKMS/pkg/models"
	"github.com/Mario-Kamel/EKMS/pkg/service"

	"github.com/gorilla/mux"
)

type AssignmentController struct {
	svc *service.AssignmentService
}

func NewAssignmentController(svc *service.AssignmentService) *AssignmentController {
	return &AssignmentController{
		svc: svc,
	}
}

func (c *AssignmentController) GetAllAssignments(w http.ResponseWriter, r *http.Request) {
	assignments, err := c.svc.GetAllAssignments(context.Background())
	if err != nil {
		fmt.Printf("Error while getting all assignments: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assignments)
}

func (c *AssignmentController) GetAssignmentById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	assignment, err := c.svc.GetAssignmentById(context.Background(), id)
	if err != nil {
		var IDErr *cerrors.InvalidIDError
		if errors.As(err, &IDErr) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assignment)
}

func (c *AssignmentController) CreateAssignment(w http.ResponseWriter, r *http.Request) {
	var assignment models.Assignment
	err := json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {
		fmt.Printf("Error while decoding assignment: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := c.svc.CreateAssignment(context.Background(), assignment)
	if err != nil {
		fmt.Printf("Error while creating assignment: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (c *AssignmentController) UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var assignment models.Assignment
	err := json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {
		fmt.Printf("Error while decoding assignment: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := c.svc.UpdateAssignment(context.Background(), id, assignment)
	if err != nil {
		fmt.Printf("Error while updating assignment: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (c *AssignmentController) DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := c.svc.DeleteAssignment(context.Background(), id)
	if err != nil {
		fmt.Printf("Error while deleting assignment: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
