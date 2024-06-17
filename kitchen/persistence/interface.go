package persistence

import "database/sql"

type Persistence interface {
	CreateOrder(recipeName, status string) (Order, error)
	GetOrders() ([]Order, error)
	GetRecipes() ([]Recipe, error)
	UpdateOrderStatus(orderId int, status string) error
	GetRecipe(recipeID int) (Recipe, error)
}

var db *sql.DB
