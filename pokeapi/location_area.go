package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var apiURL string = "https://pokeapi.co/api/v2/location-area/"

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(query_param string) (LocationAreas, error) {
	var location_areas LocationAreas
	fmt.Printf("LOG: apiURL + query_param = %s%s\n", apiURL, query_param)
	res, err := http.Get(apiURL + query_param)
	if err != nil {
		return location_areas, fmt.Errorf("REQUEST FAILED: %v", err)
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&location_areas)
	if err != nil {
		return location_areas, fmt.Errorf("DECODING THE BODY FAILED:: %v", err)
	}
	return location_areas, nil
}
