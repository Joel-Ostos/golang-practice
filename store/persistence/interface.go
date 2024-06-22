package persistence

import "database/sql"

type Persistence interface {
	GetIngredientQuantity(ingredient string) (int, error)
	SetIngredientQuantity(ingredient string, quantityBought int) error
	GetStock() (Stock, error)
	GetPurchases() (Purchases, error)
}

var db *sql.DB
