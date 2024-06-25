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
		actualQuantity, err := db.GetIngredientQuantity(ingredient)
		if err != nil {
			return false, err
		}
		if actualQuantity < neededQuantity {
			wg.Add(1)
			go buyIngredient(ingredient, neededQuantity, &wg)
		}
	}
	wg.Wait()
	return true, nil
}

func buyIngredient(ingredient string, neededQuantity int, wg *sync.WaitGroup) int {
	url := "https://recruitment.alegra.com/api/farmers-market/buy?ingredient=" + ingredient
	resp, err := http.Get(url)
	defer wg.Done()

	if err != nil {
		log.Println("Error buying ingredient:", err)
		return 0
	}

	defer resp.Body.Close()

	var result map[string]int
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		log.Println("Error decoding response:", err)
		return 0
	}

	actualQuantity, err := db.GetIngredientQuantity(ingredient)

	if err != nil {
	  return 0
	}

	actualQuantity += result["quantitySold"]
	err = db.SetIngredientBought(ingredient, result["quantitySold"])

	if err != nil {
	  return 0
	}

	for actualQuantity < neededQuantity {
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error buying ingredient:", err)
			return 0
		}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Println("Error decoding response:", err)
			return 0
		}
		actualQuantity += result["quantitySold"]
		err = db.SetIngredientBought(ingredient, result["quantitySold"])
		if err != nil {
		  return 0
		}
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
	purchases, err := db.GetPurchases()

	if err != nil {
		log.Fatal(err.Error())
		return persistence.Purchases{}, err
	}

	return purchases, nil
}
