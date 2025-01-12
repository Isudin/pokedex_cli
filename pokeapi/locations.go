package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Isudin/pokedex_cli/internal/pokecache"
)

const apiUrl = "https://pokeapi.co/api/v2/"
const locationAreaEndpoint = "location-area"
const reqParams = "?offset=0&limit=20"

var cache *pokecache.Cache = &pokecache.Cache{}

func GetLocationAreas(url string) (LocationAreas, error) {
	var areas LocationAreas
	if url == "" {
		url = apiUrl + locationAreaEndpoint + reqParams
	}

	areas, err := GetAreasFromCache(url)
	if err != nil {
		return LocationAreas{}, err
	}

	if areas.Count != 0 {
		return areas, nil
	}

	body, err := Get(url)
	if err != nil {
		return LocationAreas{}, err
	}

	err = json.Unmarshal(body, &areas)
	if err != nil {
		return LocationAreas{}, err
	}

	cache.Add(url, body)
	return areas, nil
}

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

func GetAreasFromCache(url string) (LocationAreas, error) {
	var areas LocationAreas
	if !cache.Inniciated {
		newCache, err := pokecache.NewCache(pokecache.MinDuration)
		cache = newCache
		return LocationAreas{}, err
	} else if data, isCached := cache.Get(url); isCached {
		err := json.Unmarshal(data, &areas)
		return areas, err
	}

	return LocationAreas{}, nil
}

type LocationAreas struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Areas    []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
}
