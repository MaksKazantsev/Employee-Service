package app

import (
	"github.com/MaksKazantsev/mongodb/internal/routes"
	"github.com/MaksKazantsev/mongodb/internal/service"
	"github.com/MaksKazantsev/mongodb/internal/storage/mongo"
	"github.com/alserov/fuze"
)

func MustStart() {
	// Db connect
	db := mongo.MustConnect()
	repo := mongo.NewRepository(db)

	// New service

	s := service.NewService(repo)

	// Routes setup

	a := fuze.NewApp()
	routes.SetupRoutes(a, s)

	// Server run

	err := a.Run()
	if err != nil {
		panic(err)
	}
}
