package main

import (
	"kitchen/persistence"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db persistence.Persistence

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/order", createOrderHandler).Methods("POST")
	router.HandleFunc("/recipes", getRecipesHandler).Methods("GET")
	router.HandleFunc("/orders", getOrdersHandler).Methods("GET")
	db = persistence.NewPersistenceInMemory()
	log.Println("Kitchen service running on port 8080")
	http.ListenAndServe(":8080", router)
}
