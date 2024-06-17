package persistence

import (
	"log"
)

type persistenceInMemory struct {}

func NewPersistenceInMemory() Persistence {
	return persistenceInMemory{}
}

func (p persistenceInMemory) CreateOrder(recipeName, status string) (Order, error) {
    var order Order

    // Obtener el ID de la receta
    var recipeID int
    err := db.QueryRow("SELECT id FROM Recipes WHERE nombre = ?", recipeName).Scan(&recipeID)
    if err != nil {
        log.Fatal(err.Error())
	log.Printf("No funca acá 1")
        return Order{}, err
    }

    // Insertar el nuevo pedido
    _, err = db.Exec("INSERT INTO Orders (receta_id, status) VALUES (?, ?)", recipeID, status)
    if err != nil {
        log.Fatal(err.Error())
	log.Printf("No funca acá 1")
        return Order{}, err
    }

    // Obtener el ID del nuevo pedido
    err = db.QueryRow("SELECT id FROM Orders WHERE receta_id = ? ORDER BY id DESC LIMIT 1", recipeID).Scan(&order.ID)
    if err != nil {
        log.Fatal(err.Error())
	log.Printf("No funca acá 1")
        return Order{}, err
    }

    // Asignar los valores a la estructura Order
    order.Recipe = recipeName
    order.Status = status

    return order, nil
}
func (p persistenceInMemory) GetOrders() ([]Order, error) {
	result, err := db.Query("SELECT * FROM Orders")
	if err != nil {
	  log.Fatal(err.Error())
	  return []Order{}, nil
	}
	var orders []Order

	for result.Next() {
	  var tmp Order
	  var recipe_id int
	  result.Scan(&tmp.ID, &recipe_id, &tmp.Status)
	  result := db.QueryRow("SELECT nombre FROM Recipes where id = ? ", recipe_id)
	  err := result.Scan(&tmp.Recipe)
	  if err != nil {
	    log.Fatal(err.Error())
	    return []Order{}, err
	  }
	  orders = append(orders, tmp)
	}
	return orders, nil
}

func (p persistenceInMemory) GetRecipes() ([]Recipe, error) {
	result, err := db.Query("SELECT * FROM Recipes")
	if err != nil {
	  log.Fatal(err.Error())
	  return []Recipe{}, nil
	}

	var recipes []Recipe

	for result.Next() {
	  var tmp Recipe
	  result.Scan(&tmp.ID, &tmp.Name)
	  recipesIngredients, err := db.Query("SELECT ingrediente_id, cantidad FROM RecipeIngredients where receta_id = ?", tmp.ID)
	  if err != nil {
	    log.Fatal(err.Error())
	    return []Recipe{}, err
	  }
	  m := make(map[string]int)
	  tmp.Ingredients = m
	  for recipesIngredients.Next() {
	    var ingredientId, quantity int
	    var ingredientName string

	    recipesIngredients.Scan(&ingredientId, &quantity)
	    ingredientNameRow := db.QueryRow("SELECT nombre FROM Ingredients where id = ?", ingredientId)
	    err = ingredientNameRow.Scan(&ingredientName)
	    if err != nil {
		log.Fatal(err.Error())
		return []Recipe{}, err
	    }
	    m[ingredientName] = quantity
	  }
	  recipes = append(recipes, tmp)
	}
	return recipes, nil
}

func (p persistenceInMemory) GetRecipe(recipeID int) (Recipe, error) {
	var tmp Recipe
	tmp.ID = recipeID
	recipeName :=  db.QueryRow("SELECT nombre FROM Recipes where id = ?", tmp.ID)
	recipeName.Scan(&tmp.Name)

	m := make(map[string]int)
	tmp.Ingredients = m
	recipesIngredients, err := db.Query("SELECT ingrediente_id, cantidad FROM RecipeIngredients where receta_id = ?", tmp.ID)
	if err != nil {
	  log.Fatal(err.Error())
	  return Recipe{}, err
	}
	for recipesIngredients.Next() {
	  var ingredientId, quantity int
	  var ingredientName string

	  recipesIngredients.Scan(&ingredientId, &quantity)
	  ingredientNameRow := db.QueryRow("SELECT nombre FROM Ingredients where id = ?", ingredientId)
	  err = ingredientNameRow.Scan(&ingredientName)
	  if err != nil {
	    log.Fatal(err.Error())
	    return Recipe{}, err
	  }
	  m[ingredientName] = quantity
	}
	return tmp, nil
}

func (p persistenceInMemory) UpdateOrderStatus(orderId int, status string) error {
  _, err := db.Exec("UPDATE Orders SET status = ? WHERE id = ?", status, orderId)
  if err != nil {
    log.Fatal(err.Error())
    return err
  }
  return nil
}
