package main

import (
	"fmt"
)

func main() {
	breeds, err := getCatBreedsFromAPI()
	if err != nil {
		fmt.Printf("Error getting cat breeds: %v\n", err)
		return
	}

	groupedBreeds := GroupByCountry(breeds)
	sortedBreeds := SortByLength(groupedBreeds)

	if err := saveToJSON(sortedBreeds, "out.json"); err != nil {
		fmt.Printf("Ошибка сохранения в json: %v\n", err)
		return
	}

	fmt.Println("Данные успешно сохранены в out.json.")
}
