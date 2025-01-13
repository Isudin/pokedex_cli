package pokeapi

import (
	"encoding/json"

	"github.com/Isudin/pokedex_cli/internal/pokecache"
)

var pagingParams = "?offset=0&limit=20"
var cachedLocations = &pokecache.Cache{}
var cachedPokemon = &pokecache.Cache{}

func GetLocationAreas(url string) (LocationAreas, error) {
	var areas LocationAreas
	if url == "" {
		url = apiUrl + locationAreaEndpoint + pagingParams
	}

	areas, err := getLocationAreasFromCache(url)
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

func getLocationAreasFromCache(url string) (LocationAreas, error) {
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

func GetPokemonByArea(areaName string) ([]Pokemon, error) {
	pokemon, err := getPokemonByAreaFromCache(areaName)
	if err != nil {
		return nil, err
	}

	if pokemon != nil {
		return pokemon, nil
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

	cachedPokemon.Add(areaName, data)
	return getPokemonFromEncounters(area.PokemonEncounters), nil
}

func getPokemonByAreaFromCache(areaName string) ([]Pokemon, error) {
	var area LocationArea
	if !cachedPokemon.Inniciated {
		newCache, err := pokecache.NewCache(pokecache.MinDuration)
		cachedPokemon = newCache
		return nil, err
	} else if data, isCached := cachedPokemon.Get(areaName); isCached {
		err := json.Unmarshal(data, &area)
		return getPokemonFromEncounters(area.PokemonEncounters), err
	}

	return nil, nil
}

func getPokemonFromEncounters(pokemonEncounter []PokemonEncounter) []Pokemon {
	if len(pokemonEncounter) == 0 {
		return nil
	}

	pokemon := make([]Pokemon, len(pokemonEncounter))
	for i, encounter := range pokemonEncounter {
		pokemon[i] = encounter.Pokemon
	}

	return pokemon
}

type LocationAreas struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Areas    []LocationArea `json:"results"`
}

type LocationArea struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}
