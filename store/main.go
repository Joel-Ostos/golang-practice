package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ingredients", requestIngredientsHandler).Methods("POST")
	router.HandleFunc("/stock", getStock).Methods("GET")
	router.HandleFunc("/purchases", getPurchases).Methods("GET")

	log.Println("Bodega service running on port 8081")
	http.ListenAndServe(":8081", router)
}
