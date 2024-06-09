package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func createOrder(w http.ResponseWriter, r *http.Request) {
	recipe := recipes[rand.Intn(len(recipes))]
	order := Order{ID: len(orders) + 1, Recipe: recipe.Name, Status: "recibido"}
	orders = append(orders, order)

	go requestIngredients(recipe.Ingredients, order.ID)

	json.NewEncoder(w).Encode(order)
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(recipes)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
}