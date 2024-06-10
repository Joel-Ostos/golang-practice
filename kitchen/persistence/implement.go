package persistence

type persistenceInMemory struct {}

func NewPersistenceInMemory() Persistence {
	return persistenceInMemory{}
}

func (p persistenceInMemory) CreateOrder(recipeName, status string) (Order, error) {
	order := Order{ID: len(orders) + 1, Recipe: recipeName, Status: status}
	orders = append(orders, order)
	return order, nil
}

func (p persistenceInMemory) GetOrders() ([]Order, error) {
	return orders, nil
}

func (p persistenceInMemory) GetRecipes() ([]Recipe, error) {
	return recipes, nil
}

func (p persistenceInMemory) GetRecipe(recipeID int) (Recipe, error) {
	return recipes[recipeID - 1], nil
}

func (p persistenceInMemory) UpdateOrderStatus(orderId int, status string) error {
	for i, order := range orders {
		if order.ID == orderId {
			orders[i].Status = status
			break
		}
	}
	return nil
}