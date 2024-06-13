package main

import (
	"encoding/json"
	"net/http"
)

func requestIngredientsHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	var requestData askIngredients
	json.NewDecoder(r.Body).Decode(&requestData)

	ingredientsAvailable, err := requestIngredients(&requestData)

	if (err != nil) {
	  http.Error(w, "Error buying ingredients", http.StatusInternalServerError)
	  return
	}

	if ingredientsAvailable {
		for ingredient, quantity := range requestData.Ingredients {
			stock[ingredient] -= quantity
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "ingredients available"})
		return;
	} 
	json.NewEncoder(w).Encode(map[string]string{"status": "waiting for ingredients"})
}

func getStock(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stock)
}

func getPurchases(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(purchases)
}
