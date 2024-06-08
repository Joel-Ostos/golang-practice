package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Stock map[string]int

type Purchase struct {
	Ingredient   string
	QuantitySold int
}

var stock = Stock{
	"tomato":  5,
	"lemon":   5,
	"potato":  5,
	"rice":    5,
	"ketchup": 5,
	"lettuce": 5,
	"onion":   5,
	"cheese":  5,
	"meat":    5,
	"chicken": 5,
}

var purchases []Purchase

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ingredients", requestIngredients).Methods("POST")
	router.HandleFunc("/stock", getStock).Methods("GET")
	router.HandleFunc("/purchases", getPurchases).Methods("GET")

	log.Println("Bodega service running on port 8081")
	http.ListenAndServe(":8081", router)
}

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

func buyIngredient(ingredient string) int {
	url := "https://recruitment.alegra.com/api/farmers-market/buy?ingredient=" + ingredient
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error buying ingredient:", err)
		return 0
	}
	defer resp.Body.Close()

	var result map[string]int
	err_on := json.NewDecoder(resp.Body).Decode(&result)
	if err_on != nil {
	  log.Println("Error decoding response:", err)
	  return 0
	}

	for result["quantitySold"] == 0 {
	  resp, err := http.Get(url)
	  if err != nil {
	    log.Println("Error buying ingredient:", err)
	    return 0
	  }
	  defer resp.Body.Close()
	  err_on := json.NewDecoder(resp.Body).Decode(&result)
	  if err_on != nil {
	    log.Println("Error decoding response:", err)
	    return 0
	  }
	}
	return result["quantitySold"]
}

func jsonRequest(data interface{}) *bytes.Buffer {
	body, _ := json.Marshal(data)
	return bytes.NewBuffer(body)
}

