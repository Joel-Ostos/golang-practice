package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func requestIngredients(requestData *askIngredients) (bool, error){
  var wg sync.WaitGroup
  for ingredient, quantity := range requestData.Ingredients {
    if stock[ingredient] < quantity {
      wg.Add(1)
      answer := make(chan int)
      go func() {
	answer <- buyIngredient(ingredient, quantity, &wg)
      }()
      quantityBought := <-answer
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

