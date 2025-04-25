package commands

import (
	"fmt"
	"net/url"
	"os"

	"github.com/thiagovandieten/pokedex/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

var cmdMap map[string]CliCommand

type Config struct {
	Next     string
	Previous string
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

}

func CommandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(c *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range cmdMap {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func CommandMap(c *Config) error {
	var query_param string
	// fmt.Printf("Config struct: %+v", c)
	if len(c.Next) != 0 {
		query_param = constructQueryParam(c.Next)
	}
	// fmt.Printf("LOG: query_param after nlURL check = %s\n", query_param)

	la, err := pokeapi.GetLocationAreas(query_param)
	if err != nil {
		return err
	}

	if la.Previous != nil {
		c.Previous = *la.Previous
	}
	if len(la.Next) != 0 {
		c.Next = la.Next
	}

	for _, result := range la.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func CommandMapB(c *Config) error {
	// fmt.Printf("Config struct: %+v", c)
	var query_param string
	if len(c.Next) != 0 {
		query_param = constructQueryParam(c.Previous)
	} else {
		fmt.Println("You are on the first page.")
		return nil
	}

	la, err := pokeapi.GetLocationAreas(query_param)
	if err != nil {
		return err
	}

	if la.Previous != nil {
		c.Previous = *la.Previous
	}
	if len(la.Next) != 0 {
		c.Next = la.Next
	}

	for _, result := range la.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func constructQueryParam(link string) string {
	p, err := url.Parse(link)
	if err != nil {
		return ""
	}

	query_param := "?" + p.RawQuery
	return query_param
}

func ExecuteCommand(name string, args *Config) error {
	if cmd, ok := cmdMap[name]; ok {
		return cmd.Callback(args)
	}
	return fmt.Errorf("unknown command: %s", name)
}
