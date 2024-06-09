package main

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