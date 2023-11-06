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

type PersonController struct {
	svc *service.PersonService
}

func NewPersonController(svc *service.PersonService) *PersonController {
	return &PersonController{
		svc: svc,
	}
}

func (c *PersonController) GetAllPersons(w http.ResponseWriter, r *http.Request) {
	persons, err := c.svc.GetAllPersons(context.Background())
	if err != nil {
		fmt.Printf("Error while getting all persons: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func (c *PersonController) GetPersonById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	person, err := c.svc.GetPersonById(context.Background(), id)
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
	json.NewEncoder(w).Encode(person)
}

func (c *PersonController) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		fmt.Printf("Error while decoding person: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	p, err := c.svc.CreatePerson(context.Background(), person)
	if err != nil {
		fmt.Printf("Error while creating person: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (c *PersonController) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		fmt.Printf("Error while decoding person: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	p, err := c.svc.UpdatePerson(context.Background(), id, person)
	if err != nil {
		fmt.Printf("Error while updating person: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (c *PersonController) DeletePerson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := c.svc.DeletePerson(context.Background(), id)
	if err != nil {
		fmt.Printf("Error while deleting person: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
