package pokeapi

import (
	"encoding/json"

	"github.com/Isudin/pokedex_cli/internal/pokecache"
)

var cachedPokemon = &pokecache.Cache{}

func GetPokemonByName(name string) (Pokemon, error) {
	pokemon, err := getPokemonFromCache(name)
	if err != nil {
		return Pokemon{}, err
	}

	if pokemon.Name != "" {
		return pokemon, nil
	}

	url := apiUrl + pokemonEndpoint + "/" + name

	data, err := Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	var pok Pokemon
	err = json.Unmarshal(data, &pok)
	if err != nil {
		return Pokemon{}, err
	}

	cachedPokemon.Add(name, data)
	return Pokemon{}, nil
}

func getPokemonFromCache(name string) (Pokemon, error) {
	var pokemon Pokemon
	if !cachedPokemon.Innitiated {
		newCache, err := pokecache.NewCache(pokecache.MinDuration)
		cachedPokemon = newCache
		return Pokemon{}, err
	} else if data, isCached := cachedPokemon.Get(name); isCached {
		err := json.Unmarshal(data, &pokemon)
		return pokemon, err
	}

	return Pokemon{}, nil
}
