package main

import (
	"encoding/json"
	"net/http"
)

type CatBreed struct {
	Breed  string `json:"breed"`
	Origin string `json:"country"`
}

type CatBreedsResponse struct {
	Data []CatBreed `json:"data"`
}

func getCatBreedsFromAPI() ([]CatBreed, error) {
	url := "https://catfact.ninja/breeds"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var breedsResponse CatBreedsResponse
	if err := json.NewDecoder(response.Body).Decode(&breedsResponse); err != nil {
		return nil, err
	}

	return breedsResponse.Data, nil
}
