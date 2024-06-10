package persistence

type Order struct {
	ID     int
	Recipe string
	Status string
}

type Recipe struct {
	ID          int
	Name        string
	Ingredients map[string]int
}

type Ingredient struct {
	ID   int
	Name string
}
