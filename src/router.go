package main

import (
	"log"
	"net/http"

	"github.com/Tech3790/stock-saver-backend/src/functions"
	"github.com/gorilla/mux"
)

// InitRouter starts the server and makes routes available
func InitRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", functions.ListUsers).Methods("GET")
	router.HandleFunc("/users/{id}", functions.GetUser).Methods("GET")
	router.HandleFunc("/users", functions.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", functions.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", functions.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
