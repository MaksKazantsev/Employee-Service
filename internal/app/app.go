package app

import (
	"github.com/MaksKazantsev/mongodb/internal/log"
	"github.com/MaksKazantsev/mongodb/internal/routes"
	"github.com/MaksKazantsev/mongodb/internal/service"
	"github.com/MaksKazantsev/mongodb/internal/storage/mongo"
	"github.com/alserov/fuze"
)

func MustStart() {
	// Logger

	l := log.MustSetup()

	// Db connect
	db := mongo.MustConnect()
	repo := mongo.NewRepository(db)

	// New service

	s := service.NewService(repo)
	l.Info("New service created!")

	// Routes setup

	a := fuze.NewApp()
	routes.SetupRoutes(a, s)
	l.Info("Routes generated!")

	// Server run

	err := a.Run()
	l.Info("Server started!")
	if err != nil {
		panic(err)
	}
}
