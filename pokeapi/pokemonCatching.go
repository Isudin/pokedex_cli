package pokeapi

import (
	"github.com/Isudin/pokedex_cli/internal/pokecache"
)

var cachedPokemon = &pokecache.Cache{}

func GetPokemonByName(name string) (Pokemon, error) {
	//Try get pokemon from cache
	pokemon, err := getPokemonFromCache()
	if err != nil {
		return Pokemon{}, err
	}

	if pokemon.Name != "" {
		return pokemon, nil
	}

	//Prepare URL

	//Do Get request from API
	// Get()

	//Unmarshal data

	//Save data to cache

	//return pokemon

	return Pokemon{}, nil
}

func getPokemonFromCache(name string) (Pokemon, error) {
	return Pokemon{}, nil
}
