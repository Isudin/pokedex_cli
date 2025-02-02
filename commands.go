package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/Isudin/pokedex_cli/pokeapi"
)

func commandExit(_ *pokeapi.LocationAreas, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *pokeapi.LocationAreas, _ []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandMap(areas *pokeapi.LocationAreas, _ []string) error {
	return mapLocations(areas, true)
}

func commandMapb(areas *pokeapi.LocationAreas, _ []string) error {
	return mapLocations(areas, false)
}

func mapLocations(areas *pokeapi.LocationAreas, isNext bool) error {
	url := ""

	if areas.Count != 0 {
		if isNext {
			if areas.Next == nil {
				fmt.Println("you're on the last page")
				return nil
			}
			url = *areas.Next
		} else {
			if areas.Previous == nil {
				fmt.Println("you're on the first page")
				return nil
			}
			url = *areas.Previous
		}
	}

	locationAreas, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas.Areas {
		fmt.Println(locationArea.Name)
	}

	*areas = locationAreas
	return nil
}

func commandExplore(_ *pokeapi.LocationAreas, parameters []string) error {
	if parameters == nil {
		return fmt.Errorf("no parameters found")
	}

	area := parameters[0]
	fmt.Printf("Exploring area %v...\n", area)
	pokemon, err := pokeapi.GetPokemonByArea(area)
	if err != nil {
		return err
	}

	if len(pokemon) == 0 {
		fmt.Println("No Pokemon found in this location")
	} else {
		fmt.Println("Found Pokemon:")
		for _, pok := range pokemon {
			fmt.Printf(" - %v\n", pok.Name)
		}
	}

	return nil
}

func commandCatch(_ *pokeapi.LocationAreas, parameters []string) error {
	if parameters == nil {
		return fmt.Errorf("no parameters found")
	}

	name := parameters[0]
	pokemon, err := pokeapi.GetPokemonByName(name)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")
	if rand.Float64() <= 0.2 {
		fmt.Println(pokemon.Name + " caught!")
		pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name + " has escaped!")
	}

	return nil
}

func commandInspect(_ *pokeapi.LocationAreas, parameters []string) error {
	if parameters == nil {
		return fmt.Errorf("no parameters found")
	}

	name := parameters[0]
	pokemon, exists := pokedex[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%v: %v\n", stat.StatNameObj.Name, stat.Value)
	}

	fmt.Printf("Types: \n")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("- %v\n", pokeType.PokemonType.Name)
	}

	return nil
}

func commandPokedex(_ *pokeapi.LocationAreas, _ []string) error {
	if len(pokedex) == 0 {
		fmt.Println("Your pokedex is empty")
		return nil
	}

	fmt.Println("Your pokedex:")
	for _, pok := range pokedex {
		fmt.Println("- " + pok.Name)
	}

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Get some help",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displaying names of the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displaying names of the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all pokemons found in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all of pokemon you've caught",
			callback:    commandPokedex,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.LocationAreas, []string) error
}
