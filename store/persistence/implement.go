package persistence

import (
	"log"
)

type persistenceInMemory struct{}

func NewPersistenceInMemory() Persistence {
	return persistenceInMemory{}
}

func (p persistenceInMemory) GetIngredientQuantity(ingredient string) (int, error) {
	var quantity int
	statementId := db.QueryRow("SELECT cantidad FROM Stock WHERE nombre = ?", ingredient)
	err := statementId.Scan(&quantity)
	if err != nil {
		log.Fatal(err.Error())
		return -1, err
	}
	return quantity, nil
}

// TODO Arreglar para no agregar valores nulos a la tabla de purchases cuando se retira un ingrediente

func (p persistenceInMemory) SetIngredientQuantity(ingredient string, quantityBought int) error {
	var quantity int
	statementId := db.QueryRow("SELECT cantidad FROM Stock WHERE nombre = ?", ingredient)

	err := statementId.Scan(&quantity)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	_, err = db.Exec("UPDATE Stock SET cantidad = ? WHERE nombre = ?", quantity+quantityBought, ingredient)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	if quantityBought > 0 {
		_, err = db.Exec("INSERT INTO Purchases(ingrediente_id,cantidad_comprada,status)", statementId, quantityBought, "OK")
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
		return nil
	}
	_, err = db.Exec("INSERT INTO Purchases(ingrediente_id,cantidad_comprada,status)", statementId, quantityBought, "FAIL")
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (p persistenceInMemory) GetStock() (Stock, error) {
	var st Stock
	rows, err := db.Query("SELECT ingrediente_id, cantidad FROM Stock")

	if err != nil {
		return Stock{}, err
	}

	for rows.Next() {
		var ingredientId, quantity int
		var ingredientName string
		err := rows.Scan(&ingredientId, &quantity)
		if err != nil {
			log.Fatal(err.Error())
			return Stock{}, err
		}
		statement := db.QueryRow("SELECT nombre FROM Ingredients WHERE id = ?", ingredientId)
		err = statement.Scan(&ingredientName)
		if err != nil {
			log.Fatal(err.Error())
			return Stock{}, err
		}
		st[ingredientName] = quantity
	}
	return st, nil
}

func (p persistenceInMemory) GetPurchases() (Purchases, error) {
	var pr Purchases
	rows, err := db.Query("SELECT ingrediente_id, cantidad_comprada, status FROM Stock")

	if err != nil {
		return Purchases{}, err
	}

	for rows.Next() {
		var actualPurchase Purchase
		var ingredientId int
		err := rows.Scan(&ingredientId, &actualPurchase.QuantitySold, &actualPurchase.Status)
		if err != nil {
			log.Fatal(err.Error())
			return Purchases{}, err
		}
		statement := db.QueryRow("SELECT nombre FROM Ingredients WHERE id = ?", ingredientId)
		err = statement.Scan(&actualPurchase.Ingredient)
		if err != nil {
			log.Fatal(err.Error())
			return Purchases{}, err
		}
		pr = append(pr, actualPurchase)
	}
	return pr, nil
}
