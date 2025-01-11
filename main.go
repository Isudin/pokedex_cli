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

		text := scanner.Text()
		command := cleanInput(text)

	}
}

func cleanInput(text string) []string {
	lowerString := strings.ToLower(text)
	return strings.Fields(lowerString)
}
