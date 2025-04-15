package main

import (
	"errors"
	"fmt"
	"strings"
)

func cleanInput(text string) ([]string, error) {
	if len(text) == 0 {
		return []string{}, errors.New("input is empty")
	}
	return strings.Split(strings.ToLower(strings.TrimSpace(text)), " "), nil
}

func main() {
	fmt.Println("Hello, World!")
}
