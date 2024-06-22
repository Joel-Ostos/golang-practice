package persistence

type Stock map[string]int

type Purchase struct {
	Ingredient   string
	QuantitySold int
	Status string
}
type Purchases []Purchase
