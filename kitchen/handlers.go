package main

import (
	"encoding/json"
	"net/http"
)

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	order, err := createOrder()
	if err != nil {
		http.Error(w, "Error creating order", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(order)
}

func getRecipesHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	recipes, err := getRecipes()
	if err != nil {
		http.Error(w, "Error getting recipes", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(recipes)
}

func getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	orders, err := getOrders()
	if err != nil {
		http.Error(w, "Error getting orders", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}