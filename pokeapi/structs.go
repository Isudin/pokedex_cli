package pokeapi

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
	Name   string         `json:"name"`
	Height int            `json:"height"`
	Weight int            `json:"weight"`
	Stats  []PokemonStat  `json:"stats"`
	Types  []PokemonTypes `json:"types"`
}

type PokemonStat struct {
	Value       int  `json:"base_stat"`
	StatNameObj Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonTypes struct {
	PokemonType Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}
