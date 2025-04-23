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
	Callback    func() error
}

var cmdMap map[string]CliCommand
var nextLocationURL, prevLocationURL string

func init() {
	cmdMap = make(map[string]CliCommand)

	cmdMap["exit"] = CliCommand{
		Name:        "exit",
		Description: "Exit the pokedex",
		Callback:    CommandExit,
	}

	cmdMap["map"] = CliCommand{
		Name:        "map",
		Description: "Shows you 20 location areas in Pokemon",
		Callback:    CommandMap,
	}

	cmdMap["help"] = CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    CommandHelp,
	}
}

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range cmdMap {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func CommandMap() error {
	var query_param string
	fmt.Printf("LOG: query_param = %s\n", query_param)
	if len(nextLocationURL) != 0 {
		p, err := url.Parse(nextLocationURL)
		if err != nil {
			return err
		}

		query_param = "?" + p.RawQuery
	}
	fmt.Printf("LOG: query_param after nlURL check = %s\n", query_param)

	la, err := pokeapi.GetLocationAreas(query_param)
	if err != nil {
		return err
	}

	if len(la.Previous) != 0 {
		prevLocationURL = la.Previous
	}
	if len(la.Next) != 0 {
		nextLocationURL = la.Next
	}

	for _, result := range la.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func ExecuteCommand(name string) error {
	if cmd, ok := cmdMap[name]; ok {
		return cmd.Callback()
	}
	return fmt.Errorf("unknown command: %s", name)
}
