package main

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