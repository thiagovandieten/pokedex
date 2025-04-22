package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/thiagovandieten/pokedex/commands"
)

// type cliCommand struct {
// 	name        string
// 	description string
// 	callback    func() error
// }

// var currentCommands map[string]CliCommand

func cleanInput(text string) ([]string, error) {
	if len(text) == 0 {
		return []string{}, errors.New("input is empty")
	}
	return strings.Split(strings.ToLower(strings.TrimSpace(text)), " "), nil
}

// func commandExit() error {
// 	fmt.Println("Closing the Pokedex... Goodbye!")
// 	os.Exit(0)
// 	return nil
// }

// func commandMap() error {
// 	return nil
// }

// func commandHelp() error {
// 	fmt.Println("Welcome to the Pokedex!")
// 	fmt.Println("Usage:")
// 	for _, cmd := range commands {
// 		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
// 	}
// 	return nil
// }

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
		fmt.Printf("Your command is: %s\n", cleanedLine[0])

		if len(cleanedLine) > 0 {
			fmt.Printf("Your command is: %s\n", cleanedLine[0])
			if err := commands.ExecuteCommand(cleanedLine[0]); err != nil {
				fmt.Println(err)
			}
		}

	}
}
