package storage

import (
	"context"
	"github.com/MaksKazantsev/mongodb/internal/models"
)

type Storage interface {
	Add(ctx context.Context, e *models.Employee) error
	Get(ctx context.Context, id int) (models.Employee, error)
	GetAll(ctx context.Context) ([]models.Employee, error)
	Delete(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) (error, int64)
	Update(ctx context.Context, id int, e models.Employee) error

	CreateGroup(ctx context.Context, g *models.EmployeeGroup) error
	GetGroup(ctx context.Context, id int) (*models.EmployeeGroup, error)
	DeleteGroup(ctx context.Context, id int) error
	AddEmployeeToGroup(ctx context.Context, e models.Employee, g *models.EmployeeGroup) error
	DeleteEmployeeFromGroup(ctx context.Context, e models.Employee, g *models.EmployeeGroup) error
}
