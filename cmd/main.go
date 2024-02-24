package main

import (
	"github.com/MaksKazantsev/mongodb/internal/service"
	"github.com/MaksKazantsev/mongodb/internal/storage/mongo"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db := mongo.MustConnect()
	repo := mongo.NewRepository(db)
	s := service.NewService(repo)
	r := mux.NewRouter()
	r.HandleFunc("/employee/", s.CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{ID}", s.GetEmployee).Methods("GET")
	r.HandleFunc("/employee/{ID}", s.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/employee/{ID}", s.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/", s.GetAllEmployee).Methods("GET")
	r.HandleFunc("/employee/", s.DeleteAllEmployee).Methods("DELETE")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
