package main

func GroupByCountry(breeds []CatBreed) map[string][]string {
	countryDict := make(map[string][]string)
	for _, breed := range breeds {
		countryDict[breed.Origin] = append(countryDict[breed.Origin], breed.Breed)
	}
	return countryDict
}
