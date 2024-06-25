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
    var recipeID int

    err := db.QueryRow(`
    SELECT id 
    FROM Recipes 
    WHERE nombre = ?`, recipeName).Scan(&recipeID)

    if err != nil {
        log.Fatal(err.Error())
        return Order{}, err
    }

    _, err = db.Exec(`
    INSERT INTO Orders (receta_id, status) 
    VALUES (?, ?)`, recipeID, status)

    if err != nil {
        return Order{}, err
    }

    err = db.QueryRow(`
    SELECT id 
    FROM Orders 
    WHERE receta_id = ? 
    ORDER BY id DESC LIMIT 1`, recipeID).Scan(&order.ID)

    if err != nil {
        log.Fatal(err.Error())
        return Order{}, err
    }

    order.Recipe = recipeName
    order.Status = status

    return order, nil
}
func (p persistenceInMemory) GetOrders() ([]Order, error) {
	var orders []Order

	result, err := db.Query(`SELECT * FROM Orders`)
	if err != nil {
	  return []Order{}, nil
	}

	for result.Next() {
	  var actualOrder Order
	  var recipe_id int
	  result.Scan(&actualOrder.ID, &recipe_id, &actualOrder.Status)
	  err = db.QueryRow(`
	  SELECT nombre 
	  FROM Recipes 
	  WHERE id = ?`, recipe_id).Scan(&actualOrder.Recipe)

	  if err != nil {
	    return []Order{}, err
	  }

	  orders = append(orders, actualOrder)
	}
	return orders, nil
}

func (p persistenceInMemory) GetRecipes() ([]Recipe, error) {
	var recipes []Recipe

	result, err := db.Query(`SELECT * FROM Recipes`)
	if err != nil {
	  log.Fatal(err.Error())
	  return []Recipe{}, nil
	}


	for result.Next() {
	  var actualRecipe Recipe
	  result.Scan(&actualRecipe.ID, &actualRecipe.Name)
	  recipesIngredients, err := db.Query(`
	  SELECT ingrediente_id, cantidad 
	  FROM RecipeIngredients 
	  WHERE receta_id = ?`, actualRecipe.ID)

	  if err != nil {
	    return []Recipe{}, err
	  }

	  m := make(map[string]int)
	  actualRecipe.Ingredients = m

	  for recipesIngredients.Next() {
	    var ingredientId, quantity int
	    var ingredientName string

	    recipesIngredients.Scan(&ingredientId, &quantity)
	    err = db.QueryRow(`
	    SELECT nombre 
	    FROM Ingredients 
	    WHERE id = ?`, ingredientId).Scan(&ingredientName)

	    if err != nil {
		return []Recipe{}, err
	    }

	    m[ingredientName] = quantity
	  }

	  recipes = append(recipes, actualRecipe)
	}
	return recipes, nil
}

func (p persistenceInMemory) GetRecipe(recipeID int) (Recipe, error) {
	var actualRecipe Recipe
	actualRecipe.ID = recipeID

	err :=  db.QueryRow(`
	SELECT nombre 
	FROM Recipes 
	WHERE id = ?`, actualRecipe.ID).Scan(&actualRecipe.Name)

	if err != nil {
	  return Recipe{}, err
	}

	m := make(map[string]int)
	actualRecipe.Ingredients = m

	recipesIngredients, err := db.Query(`
	SELECT ingrediente_id, cantidad 
	FROM RecipeIngredients 
	WHERE receta_id = ?`, actualRecipe.ID)

	if err != nil {
	  return Recipe{}, err
	}

	for recipesIngredients.Next() {
	  var ingredientId, quantity int
	  var ingredientName string
	  recipesIngredients.Scan(&ingredientId, &quantity)

	  err = db.QueryRow(`
	  SELECT nombre 
	  FROM Ingredients 
	  WHERE id = ?`, ingredientId).Scan(&ingredientName)

	  if err != nil {
	    return Recipe{}, err
	  }

	  m[ingredientName] = quantity
	}
	return actualRecipe, nil
}

func (p persistenceInMemory) UpdateOrderStatus(orderId int, status string) error {
  _, err := db.Exec(`
  UPDATE Orders 
  SET status = ? 
  WHERE id = ?`, status, orderId)

  if err != nil {
    return err
  }
  return nil
}
