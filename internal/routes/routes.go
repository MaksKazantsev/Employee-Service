package routes

import (
	"github.com/MaksKazantsev/mongodb/internal/service"
	"github.com/alserov/fuze"
)

func SetupRoutes(a *fuze.App, s *service.Service) {
	a.GET("/employee/", s.GetAllEmployee)
	a.POST("/employee/{id}", s.CreateEmployee)
	a.GET("/employee/{id}", s.GetEmployee)
	a.DELETE("/employee/{id}", s.DeleteEmployee)
	a.PUT("/employee/{id}", s.UpdateEmployee)
	a.DELETE("/employee/", s.DeleteAllEmployee)

	a.POST("/group/", s.NewGroup)
	a.DELETE("/group/{id}", s.DeleteGroup)
	a.PUT("/group/{id}/{employeeId}", s.AddEmployeeToGroup)
	a.DELETE("/group/{id}/{employeeId}", s.DeleteEmployeeFromGroup)
	a.GET("/group/{id}", s.GetGroup)
}
