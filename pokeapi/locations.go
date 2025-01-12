package pokeapi

import (
	"encoding/json"

	"github.com/Isudin/pokedex_cli/internal/pokecache"
)

var pagingParams = "?offset=0&limit=20"
var cachedLocations = &pokecache.Cache{}
var cachedPokemons = &pokecache.Cache{}

func GetLocationAreas(url string) (LocationAreas, error) {
	var areas LocationAreas
	if url == "" {
		url = apiUrl + locationAreaEndpoint + pagingParams
	}

	areas, err := GetLocationAreasFromCache(url)
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

	cachedLocations.Add(url, body)
	return areas, nil
}

func GetLocationAreasFromCache(url string) (LocationAreas, error) {
	var areas LocationAreas
	if !cachedLocations.Inniciated {
		newCache, err := pokecache.NewCache(pokecache.MinDuration)
		cachedLocations = newCache
		return LocationAreas{}, err
	} else if data, isCached := cachedLocations.Get(url); isCached {
		err := json.Unmarshal(data, &areas)
		return areas, err
	}

	return LocationAreas{}, nil
}

func GetPokemonsByArea(areaName string) ([]Pokemon, error) {
	pokemons, err := GetPokemonsByAreaFromCache(areaName)
	if err != nil {
		return nil, err
	}

	if pokemons != nil {
		return pokemons, nil
	}

	url := apiUrl + locationAreaEndpoint + "/" + areaName
	data, err := Get(url)
	if err != nil {
		return nil, err
	}

	var area LocationArea
	err = json.Unmarshal(data, &area)
	if err != nil {
		return nil, err
	}

	return area.Pokemons, nil
}

func GetPokemonsByAreaFromCache(areaName string) ([]Pokemon, error) {
	var area LocationArea
	if !cachedPokemons.Inniciated {
		newCache, err := pokecache.NewCache(pokecache.MinDuration)
		cachedPokemons = newCache
		return nil, err
	} else if data, isCached := cachedPokemons.Get(areaName); isCached {
		err := json.Unmarshal(data, &area)
		return area.Pokemons, err
	}

	return nil, nil
}

type LocationAreas struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Areas    []LocationArea `json:"results"`
}

type LocationArea struct {
	Name     string    `json:"name"`
	Pokemons []Pokemon `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string `json:"name"`
}
