package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"store/persistence"
)

var db persistence.Persistence

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ingredients", requestIngredientsHandler).Methods("POST")
	router.HandleFunc("/stock", getStockHandler).Methods("GET")
	router.HandleFunc("/purchases", getPurchasesHandler).Methods("GET")
	db = persistence.NewPersistenceInMemory()
	log.Println("Bodega service running on port 8081")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
