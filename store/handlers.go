package main

import (
	"encoding/json"
	"net/http"
)

func requestIngredients(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Ingredients map[string]int `json:"ingredients"`
		OrderID     int            `json:"orderID"`
	}
	_ = json.NewDecoder(r.Body).Decode(&requestData)

	allIngredientsAvailable := true

	for ingredient, quantity := range requestData.Ingredients {
		if stock[ingredient] < quantity {
			allIngredientsAvailable = false
			quantityBought := buyIngredient(ingredient)
			stock[ingredient] += quantityBought
			purchases = append(purchases, Purchase{Ingredient: ingredient, QuantitySold: quantityBought})
		}
		if stock[ingredient] < quantity {
			allIngredientsAvailable = false
			break
		}
	}

	if allIngredientsAvailable {
		for ingredient, quantity := range requestData.Ingredients {
			stock[ingredient] -= quantity
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "ingredients available"})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"status": "waiting for ingredients"})
	}
}

func getStock(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stock)
}

func getPurchases(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(purchases)
}