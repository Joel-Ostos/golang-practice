package persistence

import (
	"log"
)

type persistenceInMemory struct{}

func NewPersistenceInMemory() Persistence {
	return persistenceInMemory{}
}

func (p persistenceInMemory) GetIngredientQuantity(ingredient string) (int, error) {
	var quantity, ingredientId int

	err := db.QueryRow(`
	SELECT id 
	FROM Ingredients 
	WHERE nombre = ?`, ingredient).Scan(&ingredientId)

	if err != nil {
		return 0, err
	}

	err = db.QueryRow(`
	SELECT cantidad 
	FROM Stock 
	WHERE ingrediente_id = ?`, ingredientId).Scan(&quantity)

	if err != nil {
		return 0, err
	}

	return quantity, nil
}

func (p persistenceInMemory) SetIngredientBought(ingredient string, quantityBought int) error {
	var ingredientId, quantity int

	err := db.QueryRow(`
	SELECT id 
	FROM Ingredients 
	WHERE nombre = ?`, ingredient).Scan(&ingredientId)

	if err != nil {
		return err
	}

	err = db.QueryRow(`
	SELECT cantidad 
	FROM Stock
	WHERE ingrediente_id = ?`, ingredientId).Scan(&quantity)

	if err != nil {
	  return err 
	}

	_, err = db.Exec(`
	UPDATE Stock 
	SET cantidad = ? 
	WHERE ingrediente_id = ?`, quantity+quantityBought, ingredientId)

	if err != nil {
		return err
	}

	if quantityBought > 0 {
		log.Println("Pasa")

		_, err = db.Exec(`
		INSERT INTO Purchases(ingrediente_id,cantidad_comprada,status) 
		VALUES(?,?,?)`, ingredientId, quantityBought, "OK")

		if err != nil {
			return err
		}
		return nil
	}

	_, err = db.Exec(`
	INSERT INTO Purchases(ingrediente_id,cantidad_comprada,status)
	VALUES(?,?,?)`, ingredientId, quantityBought, "FAIL")

	if err != nil {
		return err
	}

	return nil
}

func (p persistenceInMemory) SetIngredientSold(ingredient string, quantitySold int) error {
	var quantity, ingredientId int

	err := db.QueryRow("SELECT id FROM Ingredients WHERE nombre = ?", ingredient).Scan(&ingredientId)

	if err != nil {
		return err
	}

	err = db.QueryRow(`
	SELECT cantidad 
	FROM Stock 
	WHERE ingrediente_id = ?`, ingredientId).Scan(&quantity)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
	UPDATE Stock 
	SET cantidad = ?
	WHERE ingrediente_id = ?`, quantity-quantitySold, ingredientId)

	if err != nil {
		return err
	}
	return nil
}

func (p persistenceInMemory) GetStock() (Stock, error) {
	stock := make(Stock)
	rows, err := db.Query(`
	SELECT ingrediente_id, cantidad 
	FROM Stock`)

	if err != nil {
		return Stock{}, err
	}

	for rows.Next() {
		var ingredientId, quantity int
		var ingredientName string
		err := rows.Scan(&ingredientId, &quantity)

		if err != nil {
			return Stock{}, err
		}

		err = db.QueryRow(`
		SELECT nombre 
		FROM Ingredients 
		WHERE id = ?`, ingredientId).Scan(&ingredientName)

		if err != nil {
			return Stock{}, err
		}
		stock[ingredientName] = quantity
	}

	return stock, nil
}

func (p persistenceInMemory) GetPurchases() (Purchases, error) {
	var purchases Purchases
	rows, err := db.Query(`
	SELECT ingrediente_id, cantidad_comprada, status
	FROM Purchases`)

	if err != nil {
		return Purchases{}, err
	}

	for rows.Next() {
		var actualPurchase Purchase
		var ingredientId int
		err := rows.Scan(&ingredientId, &actualPurchase.QuantitySold, &actualPurchase.Status)

		if err != nil {
			return Purchases{}, err
		}

		err = db.QueryRow(`
		SELECT nombre 
		FROM Ingredients 
		WHERE id = ?`, ingredientId).Scan(&actualPurchase.Ingredient)

		if err != nil {
			return Purchases{}, err
		}

		purchases = append(purchases, actualPurchase)
	}
	return purchases, nil
}
