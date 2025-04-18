package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) ([]string, error) {
	if len(text) == 0 {
		return []string{}, errors.New("input is empty")
	}
	return strings.Split(strings.ToLower(strings.TrimSpace(text)), " "), nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex...goodbye!")
	os.Exit(0)
	return nil
}

var commands map[string]cliCommand = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedLine, err := cleanInput(scanner.Text())
		if err != nil {
			fmt.Println("no input")
			os.Exit(1)
		}
		fmt.Printf("Your command was: %s\n", cleanedLine[0])

		if c, ok := commands[cleanedLine[0]]; ok {
			c.callback()
		}

	}
}
