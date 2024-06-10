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
	// toda esta logica debe ir en funciones a parte, los handlers deben ser lo mas simples posibles y manejar unicamente la logica de http

	allIngredientsAvailable := true

	// comprar ingredientes con concurrencia, las peticiones http toman tiempo y hacerlas de manera lineal puede ser ineficiente
	for ingredient, quantity := range requestData.Ingredients {
		// wg.add(1)
		// buyIngredient(ingredient, wg)
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
	// wg.wait()
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