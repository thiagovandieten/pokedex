package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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
		fmt.Printf("Your command was: %s\n", cleanedLine[0])
	}
}
