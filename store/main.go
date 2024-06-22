package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"store/persistence"
)

var db persistence.Persistence
var xd persistence.Persistence

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ingredients", requestIngredientsHandler).Methods("POST")
	router.HandleFunc("/stock", getStock).Methods("GET")
	router.HandleFunc("/purchases", getPurchases).Methods("GET")
	db = persistence.NewPersistenceInMemory()
	db.getIngredientQuantity("tomato")
	xd = persistence.NewPersistenceInMemory()
	log.Println("Bodega service running on port 8081")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
