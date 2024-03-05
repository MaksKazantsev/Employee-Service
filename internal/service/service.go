package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/mongodb/internal/models"
	"github.com/MaksKazantsev/mongodb/internal/storage"
	"github.com/alserov/fuze"
	"math/rand"
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
	w.Header().Set("Content-Type", "application/json")

	var employee models.Employee
	if err := ctx.Decode(&employee); err != nil {
		_ = fmt.Errorf("failed to decode: %w", err)
	}
	employee.ID = rand.Intn(100)

	if err := s.storage.Add(context.Background(), &employee); err != nil {
		_ = fmt.Errorf("repository create error: %w", err)
	}

	if err := ctx.SendValue(&employee, 200); err != nil {
		_ = fmt.Errorf("failed to send value: %w", err)
	}

}

func (s *Service) GetEmployee(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")

	params := ctx.Parameters["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		_ = fmt.Errorf("failed to convert id: %w", err)
	}

	employee, err := s.storage.Get(context.Background(), id)
	if err != nil {
		_ = fmt.Errorf("repository error: %w", err)
		return
	}

	err = ctx.SendValue(&employee, 200)
	if err != nil {
		_ = fmt.Errorf("failed to send value: %w", err)
	}
}

func (s *Service) DeleteEmployee(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")

	params := ctx.Parameters["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		_ = fmt.Errorf("failed to convert id: %w", err)
	}
	if err = s.storage.Delete(context.Background(), id); err != nil {
		if err != nil {
			_ = fmt.Errorf("repository delete error: %w", err)
		}
	}

	_, _ = w.Write([]byte("Employee was successfully deleted from the storage"))
}

func (s *Service) UpdateEmployee(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")

	var employee models.Employee

	if err := ctx.Decode(&employee); err != nil {
		_ = fmt.Errorf("failed to decode value: %w", err)
	}

	params := ctx.Parameters["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		_ = fmt.Errorf("failed to convert id: %w", err)
	}

	employee.ID = id
	err = ctx.SendValue(&employee, 200)
	if err != nil {
		_ = fmt.Errorf("failed to send value: %w", err)
	}
	if err = s.storage.Update(context.Background(), id, employee); err != nil {
		_ = fmt.Errorf("repository update error: %w", err)
	}
}

func (s *Service) GetAllEmployee(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")

	employee, err := s.storage.GetAll(context.Background())
	if err != nil {
		_ = fmt.Errorf("repository get all error: %w", err)
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
		_ = fmt.Errorf("repository delete all error: %w", err)
	}

	_, _ = w.Write([]byte("Succesfully deleted from the storage:"))
	if err = ctx.SendValue(count, 200); err != nil {
		_ = fmt.Errorf("failed to send: %v", err)
	}

}

func (s *Service) NewGroup(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")

	var group models.EmployeeGroup

	if err := ctx.Decode(&group); err != nil {
		_ = fmt.Errorf("failed to decode: %w", err)
	}
	group.ID = rand.Intn(100)
	group.EmployeeList = make([]models.Employee, 0)

	if err := s.storage.CreateGroup(context.Background(), &group); err != nil {
		_ = fmt.Errorf("repository create group error: %w", err)
	}

	err := ctx.SendValue(&group, 200)
	if err != nil {
		_ = fmt.Errorf("failed to send: %w", err)
	}
}

func (s *Service) DeleteGroup(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")

	params := ctx.Parameters["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		_ = fmt.Errorf("failed to convert id: %w", err)
		return
	}

	if err = s.storage.DeleteGroup(context.Background(), id); err != nil {
		_ = fmt.Errorf("repository delete group error: %w", err)
	}

	_, _ = w.Write([]byte("A group of employee was successfully deleted from the storage"))
}

func (s *Service) GetGroup(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")

	params := ctx.Parameters["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		_ = fmt.Errorf("failed to convert id: %w", err)
	}

	group, err := s.storage.GetGroup(context.Background(), id)
	if err != nil {
		_ = fmt.Errorf("repository get group error: %w", err)
		return
	}

	err = ctx.SendValue(&group, 200)
	if err != nil {
		_ = fmt.Errorf("failed to send value: %w", err)
	}
}

func (s *Service) AddEmployeeToGroup(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")

	params := ctx.Parameters["id"]
	params2 := ctx.Parameters["employeeId"]

	id, err := strconv.Atoi(params)
	if err != nil {
		_ = fmt.Errorf("failed to convert id: %w", err)
		return
	}

	userID, err := strconv.Atoi(params2)
	if err != nil {
		_ = fmt.Errorf("failed to convert user id: %w", err)
		return
	}

	e, err := s.storage.Get(context.Background(), userID)
	if err != nil {
		_ = fmt.Errorf("repository get user error: %w", err)
	}

	group, err := s.storage.GetGroup(context.Background(), id)
	if err != nil {
		_ = fmt.Errorf("repository get group error: %w", err)
	}

	if err = s.storage.AddEmployeeToGroup(context.Background(), e, group); err != nil {
		_ = fmt.Errorf("repository add user to group error: %w", err)
	}
}

func (s *Service) DeleteEmployeeFromGroup(ctx *fuze.Ctx) {
	w := ctx.Response
	w.Header().Set("Content-Type", "application/json")

	params := ctx.Parameters["id"]
	params2 := ctx.Parameters["employeeId"]
	id, err := strconv.Atoi(params)

	if err != nil {
		_ = fmt.Errorf("failed to convert id: %w", err)
		return
	}

	userID, err := strconv.Atoi(params2)
	if err != nil {
		_ = fmt.Errorf("failed to convert user id: %w", err)
		return
	}

	e, err := s.storage.Get(context.Background(), userID)
	if err != nil {
		_ = fmt.Errorf("repository get user error: %w", err)
	}

	group, err := s.storage.GetGroup(context.Background(), id)
	if err != nil {
		_ = fmt.Errorf("repository get group error: %w", err)
	}

	if err = s.storage.DeleteEmployeeFromGroup(context.Background(), e, group); err != nil {
		_ = fmt.Errorf("repository delete user from group error: %w", err)
	}

	_, _ = w.Write([]byte("Successfuly delete from the group employee:"))
	if err = ctx.SendValue(&e, 200); err != nil {
		_ = fmt.Errorf("failed to send a value: %w", err)
	}
}
