package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type Recipe struct {
	Name        string
	Ingredients map[string]int
}

type Order struct {
	ID     int
	Recipe string
	Status string
}

var recipes = []Recipe{
	{"Ensalada de Pollo", map[string]int{"lettuce": 1, "chicken": 1, "tomato": 1, "lemon": 1}},
	{"Arroz con Pollo", map[string]int{"rice": 1, "chicken": 1, "onion": 1, "tomato": 1}},
	{"Papas Fritas con Ketchup", map[string]int{"potato": 2, "ketchup": 1}},
	{"Hamburguesa", map[string]int{"meat": 1, "cheese": 1, "lettuce": 1, "tomato": 1}},
	{"Sopa de Cebolla", map[string]int{"onion": 3, "cheese": 1}},
	{"Pollo a la Parrilla con Ensalada", map[string]int{"chicken": 1, "lettuce": 1, "tomato": 1, "lemon": 1}},
}

var orders []Order

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/order", createOrder).Methods("POST")
	router.HandleFunc("/recipes", getRecipes).Methods("GET")
	router.HandleFunc("/orders", getOrders).Methods("GET")

	log.Println("Cocina service running on port 8080")
	http.ListenAndServe(":8080", router)
}

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

func jsonRequest(data interface{}) *bytes.Buffer {
	body, _ := json.Marshal(data)
	return bytes.NewBuffer(body)
}

