package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for true {
		fmt.Print("Pokedex > ")
		hasInput := scanner.Scan()
		if !hasInput {
			continue
		}

		//text := scanner.Text()
		//command := cleanInput(text)
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	lowerString := strings.ToLower(text)
	return strings.Fields(lowerString)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
