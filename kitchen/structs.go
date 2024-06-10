package main

type IngredientsRequest struct {
	Ingredients map[string]int `json:"ingredients"`
	OrderID     int            `json:"orderID"`
}
