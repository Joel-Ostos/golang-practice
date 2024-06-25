package persistence

import "database/sql"

type Persistence interface {
	GetIngredientQuantity(ingredient string) (int, error)
	SetIngredientBought(ingredient string, quantityBought int) error
	SetIngredientSold(ingredient string, quantityBought int) error
	GetStock() (Stock, error)
	GetPurchases() (Purchases, error)
}

var db *sql.DB
