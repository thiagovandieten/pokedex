package commands

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var CMDMap map[string]CliCommand

func init() {
	CMDMap = make(map[string]CliCommand)

	CMDMap["exit"] = CliCommand{
		Name:        "exit",
		Description: "Exit the pokedex",
		Callback:    CommandExit,
	}

	CMDMap["map"] = CliCommand{
		Name:        "map",
		Description: "Shows you 20 location areas in Pokemon",
		Callback:    CommandMap,
	}

	CMDMap["help"] = CliCommand{
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
	for _, cmd := range CMDMap {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func CommandMap() error {
	return nil
}

func ExecuteCommand(name string) error {
	if cmd, ok := CMDMap[name]; ok {
		return cmd.Callback()
	}
	return fmt.Errorf("unknown command: %s", name)
}
