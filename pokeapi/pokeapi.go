package pokeapi

import (
	"io"
	"net/http"
)

const apiUrl = "https://pokeapi.co/api/v2/"
const locationAreaEndpoint = "location-area"
const pokemonEndpoint = "pokemon"

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
