package main

import (
	"github.com/MaksKazantsev/mongodb/internal/routes"
	"github.com/MaksKazantsev/mongodb/internal/service"
	"github.com/MaksKazantsev/mongodb/internal/storage/mongo"
	"github.com/alserov/fuze"
)

func main() {

	db := mongo.MustConnect()
	repo := mongo.NewRepository(db)
	s := service.NewService(repo)

	a := fuze.NewApp()
	routes.SetupRoutes(a, s)

	err := a.Run()
	if err != nil {
		panic(err)
	}
}
