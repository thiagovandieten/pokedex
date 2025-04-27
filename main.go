package main

import (
	"time"

	"github.com/thiagovandieten/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	// pokeCache := pokecache.NewCache(time.Second*15, []byte{})
	cfg := &Config{
		pokeapiClient: pokeClient,
		// pokeCache:     pokeCache,
	}
	startRepl(cfg)
}
