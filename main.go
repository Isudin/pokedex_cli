package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Isudin/pokedex_cli/pokeapi"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	commands := getCommands()
	areas := &pokeapi.LocationAreas{}

	for {
		fmt.Print("Pokedex > ")
		hasInput := scanner.Scan()
		if !hasInput {
			continue
		}

		text := scanner.Text()
		commandNames := cleanInput(text)
		if len(commandNames) == 0 {
			continue
		}
		command := commands[commandNames[0]]
		if command.name == "" || command.callback == nil {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(areas)
		if err != nil {
			fmt.Printf("%v", err.Error())
		}
	}
}

func cleanInput(text string) []string {
	lowerString := strings.ToLower(text)
	return strings.Fields(lowerString)
}

func commandExit(_ *pokeapi.LocationAreas) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *pokeapi.LocationAreas) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandMap(areas *pokeapi.LocationAreas) error {
	return mapLocations(areas, true)
}

func commandMapb(areas *pokeapi.LocationAreas) error {
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
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.LocationAreas) error
}
