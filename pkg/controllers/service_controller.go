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

type ServiceController struct {
	svc *service.ServiceService
}

func NewServiceController(svc *service.ServiceService) *ServiceController {
	return &ServiceController{
		svc: svc,
	}
}

func (c *ServiceController) GetAllServices(w http.ResponseWriter, r *http.Request) {
	services, err := c.svc.GetAllServices(context.Background())
	if err != nil {
		fmt.Printf("Error while getting all services: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func (c *ServiceController) GetServiceById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	service, err := c.svc.GetServiceById(context.Background(), id)
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
	json.NewEncoder(w).Encode(service)
}

func (c *ServiceController) CreateService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		fmt.Printf("Error while decoding service: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s, err := c.svc.CreateService(context.Background(), service)
	if err != nil {
		fmt.Printf("Error while creating service: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func (c *ServiceController) UpdateService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		fmt.Printf("Error while decoding service: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := mux.Vars(r)["id"]
	s, err := c.svc.UpdateService(context.Background(), id, service)
	if err != nil {
		var IDErr *cerrors.InvalidIDError
		if errors.As(err, &IDErr) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Printf("Error while updating service: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func (c *ServiceController) DeleteService(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := c.svc.DeleteService(context.Background(), id)
	if err != nil {
		var IDErr *cerrors.InvalidIDError
		if errors.As(err, &IDErr) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Printf("Error while deleting service: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ServiceController) AddAttendanceRecord(w http.ResponseWriter, r *http.Request) {
	var attendanceRecord models.AttendanceRecord
	err := json.NewDecoder(r.Body).Decode(&attendanceRecord)
	if err != nil {
		fmt.Printf("Error while decoding attendance record: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := mux.Vars(r)["id"]
	s, err := c.svc.AddAttendanceRecord(context.Background(), id, attendanceRecord)
	if err != nil {
		var IDErr *cerrors.InvalidIDError
		if errors.As(err, &IDErr) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Printf("Error while adding attendance record: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func (c *ServiceController) EditAttendanceRecord(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var attendanceRecord models.AttendanceRecord
	err := json.NewDecoder(r.Body).Decode(&attendanceRecord)
	if err != nil {
		fmt.Printf("Error while decoding attendance record: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	updatedService, err := c.svc.EditAttendanceRecord(context.Background(), id, attendanceRecord)
	if err != nil {
		var IDErr *cerrors.InvalidIDError
		if errors.As(err, &IDErr) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Printf("Error while editing attendance record: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedService)
}

func (c *ServiceController) DeleteAttendanceRecord(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var attendanceRecord models.AttendanceRecord
	err := json.NewDecoder(r.Body).Decode(&attendanceRecord)
	if err != nil {
		fmt.Printf("Error while decoding attendance record: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	updatedService, err := c.svc.DeleteAttendanceRecord(context.Background(), id, attendanceRecord)
	if err != nil {
		var IDErr *cerrors.InvalidIDError
		if errors.As(err, &IDErr) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Printf("Error while deleting attendance record: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedService)
}
