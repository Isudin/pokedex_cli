package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiUrl = "https://pokeapi.co/api/v2/"
const locationAreaEndpoint = "location-area"

func GetLocationAreas(url string) (LocationAreas, error) {
	if url == "" {
		url = apiUrl + locationAreaEndpoint
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	var areas LocationAreas
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areas)
	if err != nil {
		return LocationAreas{}, err
	}

	fmt.Printf("API: %v", res.Body)

	return areas, nil
}

type LocationAreas struct {
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Areas    []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
}
