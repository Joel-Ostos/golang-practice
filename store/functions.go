package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func buyIngredient(ingredient string) int {
	url := "https://recruitment.alegra.com/api/farmers-market/buy?ingredient=" + ingredient
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error buying ingredient:", err)
		return 0
	}
	defer resp.Body.Close()

	var result map[string]int
	err_on := json.NewDecoder(resp.Body).Decode(&result)
	if err_on != nil {
	  log.Println("Error decoding response:", err)
	  return 0
	}

	for result["quantitySold"] == 0 {
	  resp, err := http.Get(url)
	  if err != nil {
	    log.Println("Error buying ingredient:", err)
	    return 0
	  }
	  defer resp.Body.Close()
	  err_on := json.NewDecoder(resp.Body).Decode(&result)
	  if err_on != nil {
	    log.Println("Error decoding response:", err)
	    return 0
	  }
	}
	return result["quantitySold"]
}