package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/thiagovandieten/pokedex/commands"
)

func cleanInput(text string) ([]string, error) {
	if len(text) == 0 {
		return []string{}, errors.New("input is empty")
	}
	return strings.Split(strings.ToLower(strings.TrimSpace(text)), " "), nil
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
		fmt.Printf("Your command is: %s\n", cleanedLine[0])

		if len(cleanedLine) > 0 {
			fmt.Printf("Your command is: %s\n", cleanedLine[0])
			if err := commands.ExecuteCommand(cleanedLine[0]); err != nil {
				fmt.Println(err)
			}
		}

	}
}
