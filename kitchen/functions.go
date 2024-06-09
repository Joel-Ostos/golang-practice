package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func requestIngredients(ingredients map[string]int, orderID int) {
	url := "http://store:8081/ingredients"
	ingredientRequest := map[string]interface{}{
		"ingredients": ingredients,
		"orderID":     orderID,
	}
	resp, err := http.Post(url, "application/json", jsonRequest(ingredientRequest))
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
		updateOrderStatus(orderID, "cocinando")
		time.Sleep(5 * time.Second) 
		updateOrderStatus(orderID, "entregado")
	}
}

func updateOrderStatus(orderID int, status string) {
	for i, order := range orders {
		if order.ID == orderID {
			orders[i].Status = status
			break
		}
	}
}