package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/order", createOrder).Methods("POST")
	router.HandleFunc("/recipes", getRecipes).Methods("GET")
	router.HandleFunc("/orders", getOrders).Methods("GET")

	log.Println("Cocina service running on port 8080")
	http.ListenAndServe(":8080", router)
}


