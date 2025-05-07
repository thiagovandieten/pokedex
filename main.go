package main

import (
	"time"

	"github.com/thiagovandieten/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeDex := make(map[string]pokeapi.Pokemon)
	cfg := &Config{
		pokeapiClient: pokeClient,
		pokeDex:       &pokeDex,
	}
	startRepl(cfg)
}
