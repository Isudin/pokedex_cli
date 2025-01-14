package pokeapi

import (
	"encoding/json"

	"github.com/Isudin/pokedex_cli/internal/pokecache"
)

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
