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
		commandWords := cleanInput(text)
		if len(commandWords) == 0 {
			continue
		}
		command := commands[commandWords[0]]
		if command.name == "" || command.callback == nil {
			fmt.Println("Unknown command")
			continue
		}

		var params []string
		if len(commandWords) > 1 {
			params = commandWords[1:]
		}

		err := command.callback(areas, params)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
	}
}

func cleanInput(text string) []string {
	lowerString := strings.ToLower(text)
	return strings.Fields(lowerString)
}
