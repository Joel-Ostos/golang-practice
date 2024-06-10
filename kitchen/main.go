package main

import (
	"cocina/persistence"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var db persistence.Persistence

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/order", createOrderHandler).Methods("POST")
	router.HandleFunc("/recipes", getRecipesHandler).Methods("GET")
	router.HandleFunc("/orders", getOrdersHandler).Methods("GET")
	db = persistence.NewPersistenceInMemory()
	log.Println("Cocina service running on port 8080")
	http.ListenAndServe(":8080", router)
}


