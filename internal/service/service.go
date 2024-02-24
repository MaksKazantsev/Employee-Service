package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/MaksKazantsev/mongodb/internal/models"
	"github.com/MaksKazantsev/mongodb/internal/storage"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

type Service struct {
	storage storage.Storage
}

func NewService(s storage.Storage) *Service {
	return &Service{
		storage: s,
	}
}

func (s *Service) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		_ = errors.New("error while decoding")
	}
	employee.ID = rand.Intn(100)
	s.storage.Add(context.Background(), &employee)
	_ = json.NewEncoder(w).Encode(&employee)
}

func (s *Service) GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["ID"])
	employee, err := s.storage.Get(context.Background(), id)
	if err != nil {
		http.Error(w, "No employee founded!", http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(&employee)
}

func (s *Service) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		http.Error(w, "Failed to convert", http.StatusBadRequest)
		return
	}
	s.storage.Delete(context.Background(), id)
	_, _ = w.Write([]byte("Employee was successfully deleted from the storage"))
}

func (s *Service) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		_ = errors.New("error while decoding!")
		return
	}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		http.Error(w, "Failed to convert", http.StatusBadRequest)
		return
	}
	employee.ID = id
	_ = json.NewEncoder(w).Encode(&employee)
	s.storage.Update(context.Background(), id, employee)
}

func (s *Service) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employee, err := s.storage.GetAll(context.Background())
	if err != nil {
		_ = errors.New("Failed to get")
	}
	_ = json.NewEncoder(w).Encode(&employee)
}

func (s *Service) DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err, count := s.storage.DeleteAll(context.Background())
	if err != nil {
		_ = errors.New("failed to delete")
	}
	w.Write([]byte("Succesfully deleted from the storage:"))
	_ = json.NewEncoder(w).Encode(&count)
}
