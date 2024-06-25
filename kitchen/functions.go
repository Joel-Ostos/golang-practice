package main

import (
	"kitchen/persistence"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"time"
)

func requestIngredients(ingredients map[string]int, orderID int) {
	ingredientRequest := IngredientsRequest{
		Ingredients: ingredients,
		OrderID:     orderID,
	}
	resp, err := http.Post(storeIngredientsUrl, "application/json", jsonRequest(ingredientRequest))
	if err != nil {
		log.Println("Error requesting ingredients:", err)
		return
	}
	defer resp.Body.Close()
	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error decoding response:", err)
		return
	}

	if result["status"] == "ingredients available" {
		db.UpdateOrderStatus(orderID, cookingStatus)
		time.Sleep(5 * time.Second) 
		db.UpdateOrderStatus(orderID, deliveredStatus)
	}
}

func createOrder() (persistence.Order, error) {
	recipe, err := getRandomRecipe()

	if err != nil {
		return persistence.Order{}, err
	}

	order, err := db.CreateOrder(recipe.Name, receivedStatus)

	if err != nil {
		return persistence.Order{}, err
	}

	go requestIngredients(recipe.Ingredients, order.ID)

	return order, nil
}

func getRandomRecipe() (persistence.Recipe, error) {
	recipeID := rand.Intn(6) + 1
	recipe,  err := db.GetRecipe(recipeID)

	if err != nil {
		log.Println("Error getting recipe:", err)
		return persistence.Recipe{}, err
	}

	return recipe, nil
}

func getOrders() ([]persistence.Order, error) {
	return db.GetOrders()
}

func getRecipes() ([]persistence.Recipe, error) {
	return db.GetRecipes()
}
