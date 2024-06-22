package main

import (
	"encoding/json"
	"log"
	"net/http"
	"store/persistence"
	"sync"
)

func requestIngredients(requestData *askIngredients) (bool, error) {
	var wg sync.WaitGroup
	for ingredient, neededQuantity := range requestData.Ingredients {
		//getIngredientQuantity(ingredient)
		if stock[ingredient] < neededQuantity {
			db.getIngredientQuantity(ingredient)
			wg.Add(1)
			answer := make(chan int)
			go func() {
				answer <- buyIngredient(ingredient, neededQuantity, &wg)
			}()
			quantityBought := <-answer
			//setIngredientQuantity(ingredient, quantityBought)
			stock[ingredient] += quantityBought
			purchases = append(purchases, Purchase{Ingredient: ingredient, QuantitySold: quantityBought})
		}
	}
	wg.Wait()
	return true, nil
}

func buyIngredient(ingredient string, neededQuantity int, wg *sync.WaitGroup) int {
	url := "https://recruitment.alegra.com/api/farmers-market/buy?ingredient=" + ingredient
	resp, err := http.Get(url)
	defer wg.Done()
	defer resp.Body.Close()
	if err != nil {
		log.Println("Error buying ingredient:", err)
		return 0
	}

	var result map[string]int
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error decoding response:", err)
		return 0
	}
	actualQuantity := result["quantitySold"]
	for actualQuantity < neededQuantity {
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error buying ingredient:", err)
			return 0
		}
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Println("Error decoding response:", err)
			return 0
		}
		neededQuantity -= actualQuantity
		actualQuantity = result["quantitySold"]
	}
	return result["quantitySold"]
}

func getStock() (persistence.Stock, error) {
	stock, err := db.GetStock()
	if err != nil {
		log.Fatal(err.Error())
		return persistence.Stock{}, err
	}
	return stock, nil
}
func getPurchases() (persistence.Purchases, error) {
	purchases, err := db.GetStock()
	if err != nil {
		log.Fatal(err.Error())
		return persistence.Purchases{}, err
	}
	return purchases, nil
}
