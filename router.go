package main

import (
	"wxcloudrun-golang/service"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/login", service.Login).Methods("POST")
	return r
}
