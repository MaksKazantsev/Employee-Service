package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/MaksKazantsev/mongodb/internal/models"
	"github.com/MaksKazantsev/mongodb/internal/storage"
	"github.com/alserov/fuze"
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

func (s *Service) CreateEmployee(ctx *fuze.Ctx) {
	w := ctx.Response
	var employee models.Employee
	w.Header().Set("Content-Type", "application/json")

	if err := ctx.Decode(&employee); err != nil {
		_ = fmt.Errorf("failed to decode: %v", err)
	}
	employee.ID = rand.Intn(100)
	s.storage.Add(context.Background(), &employee)

	err := ctx.SendValue(&employee, 200)
	if err != nil {
		_ = fmt.Errorf("failed to send: %v", err)
	}

}

func (s *Service) GetEmployee(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")
	params := ctx.Parameters["id"]
	id, err := strconv.Atoi(params)
	employee, err := s.storage.Get(context.Background(), id)
	if err != nil {
		http.Error(w, "No employee founded!", http.StatusNotFound)
		return
	}
	err = ctx.SendValue(&employee, 200)
	if err != nil {
		_ = fmt.Errorf("failed to send: %v", err)
	}
}

func (s *Service) DeleteEmployee(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")
	params := ctx.Parameters["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		http.Error(w, "Failed to convert", http.StatusBadRequest)
		return
	}
	s.storage.Delete(context.Background(), id)
	_, _ = w.Write([]byte("Employee was successfully deleted from the storage"))
}

func (s *Service) UpdateEmployee(ctx *fuze.Ctx) {
	w := ctx.Response
	var employee models.Employee
	w.Header().Set("Content-Type", "application/json")
	if err := ctx.Decode(&employee); err != nil {
		_ = fmt.Errorf("failed to decode: %v", err)
	}
	params := ctx.Parameters["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		http.Error(w, "Failed to convert", http.StatusBadRequest)
		return
	}
	employee.ID = id
	err = ctx.SendValue(&employee, 200)
	if err != nil {
		_ = fmt.Errorf("failed to send: %v", err)
	}
	s.storage.Update(context.Background(), id, employee)
}

func (s *Service) GetAllEmployee(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")
	employee, err := s.storage.GetAll(context.Background())
	if err != nil {
		_ = errors.New("Failed to get")
	}
	err = ctx.SendValue(&employee, 200)
	if err != nil {
		_ = fmt.Errorf("failed to send: %v", err)
	}
}

func (s *Service) DeleteAllEmployee(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")
	err, count := s.storage.DeleteAll(context.Background())
	if err != nil {
		_ = errors.New("failed to delete")
	}
	w.Write([]byte("Succesfully deleted from the storage:"))
	err = ctx.SendValue(count, 200)
	if err != nil {
		_ = fmt.Errorf("failed to send: %v", err)
	}
}