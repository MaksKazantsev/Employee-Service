package main

import (
	"github.com/MaksKazantsev/mongodb/internal/service"
	"github.com/MaksKazantsev/mongodb/internal/storage/mongo"
	"github.com/alserov/fuze"
)

func main() {

	db := mongo.MustConnect()
	repo := mongo.NewRepository(db)
	s := service.NewService(repo)

	a := fuze.NewApp()

	a.GET("/employee/", s.GetAllEmployee)
	a.POST("/employee/{id}", s.CreateEmployee)
	a.GET("/employee/{id}", s.GetEmployee)
	a.DELETE("/employee/{id}", s.DeleteEmployee)
	a.PUT("/employee/{id}", s.UpdateEmployee)
	a.DELETE("/employee/", s.DeleteAllEmployee)

	err := a.Run()
	if err != nil {
		panic(err)
	}
}
