package main

import (
	"log"
	"net/http"
	"rest/handlers"

	"rest/models"

	"github.com/gorilla/mux"
)

func main() {

	models.SetDeafaultUser()

	mux := mux.NewRouter()

	mux.HandleFunc("/api/v1/users/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Println("Server Success :8001")
	log.Fatal(http.ListenAndServe(":8001", mux))
}
