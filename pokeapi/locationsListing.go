package pokeapi

import (
	"encoding/json"

	"github.com/Isudin/pokedex_cli/internal/pokecache"
)

var pagingParams = "?offset=0&limit=20"
var cachedLocations = &pokecache.Cache{}

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
	if !cachedLocations.Innitiated {
		newCache, err := pokecache.NewCache(pokecache.MinDuration)
		cachedLocations = newCache
		return LocationAreas{}, err
	} else if data, isCached := cachedLocations.Get(url); isCached {
		err := json.Unmarshal(data, &areas)
		return areas, err
	}

	return LocationAreas{}, nil
}
