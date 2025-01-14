package pokeapi

import (
	"encoding/json"

	"github.com/Isudin/pokedex_cli/internal/pokecache"
)

var cachedEncounters = &pokecache.Cache{}

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

	cachedEncounters.Add(areaName, data)
	return getPokemonFromEncounters(area.PokemonEncounters), nil
}

func getPokemonByAreaFromCache(areaName string) ([]Pokemon, error) {
	var area LocationArea
	if !cachedEncounters.Innitiated {
		newCache, err := pokecache.NewCache(pokecache.MinDuration)
		cachedEncounters = newCache
		return nil, err
	} else if data, isCached := cachedEncounters.Get(areaName); isCached {
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
