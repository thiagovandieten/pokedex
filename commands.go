package main

import (
	"fmt"
	"os"

	"github.com/thiagovandieten/pokedex/internal/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config, args []string) error
}

var cmdMap map[string]CliCommand

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

func init() {
	cmdMap = make(map[string]CliCommand)

	cmdMap["exit"] = CliCommand{
		Name:        "exit",
		Description: "Exit the pokedex",
		Callback:    CommandExit,
	}

	cmdMap["help"] = CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    CommandHelp,
	}

	cmdMap["map"] = CliCommand{
		Name:        "map",
		Description: "Shows you 20 location areas in Pokemon",
		Callback:    CommandMap,
	}

	cmdMap["mapb"] = CliCommand{
		Name:        "mapb",
		Description: "Show you the previous 20 locations relative to the page",
		Callback:    CommandMapB,
	}

	cmdMap["explore"] = CliCommand{
		Name:        "explore",
		Description: "Explores a area given from map and retrives all the pokemon",
		Callback:    CommandExplore,
	}

}

func CommandExit(cfg *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range cmdMap {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func CommandMap(cfg *Config, args []string) error {

	la, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	if la.Previous != nil {
		cfg.Previous = la.Previous
	}
	if la.Next != nil {
		cfg.Next = la.Next
	}

	for _, result := range la.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func CommandMapB(cfg *Config, args []string) error {
	// fmt.Printf("Config struct: %+v", c)

	la, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	if la.Previous != nil {
		cfg.Previous = la.Previous
	}
	if la.Next != nil {
		cfg.Next = la.Next
	}

	for _, result := range la.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func CommandExplore(cfg *Config, args []string) error {
	location, err := cfg.pokeapiClient.ListPokemon(args[0])
	if err != nil {
		return err
	}

	if len(location.PokemonEncounters) < 1 {
		return fmt.Errorf("no pokemon in area")
	}

	fmt.Println("Found Pokemon:")
	for _, v := range location.PokemonEncounters {
		fmt.Printf("- %s\n", v.Pokemon.Name)
	}
	return nil

}
func ExecuteCommand(name string, cfg *Config, args []string) error {
	if cmd, ok := cmdMap[name]; ok {
		return cmd.Callback(cfg, args)
	}
	return fmt.Errorf("unknown command: %s", name)
}
