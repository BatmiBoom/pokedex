package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/BatmiBoom/pokedex/cmd/config"
	"github.com/BatmiBoom/pokedex/cmd/pokeapi"
)

const baseURI = "https://pokeapi.co/api/v2/"
const locationAreasURI = baseURI + "location-area/"

type Command struct {
	Name        string
	Description string
	Callback    func(cfg *config.Config) error
}

var commands = map[string]Command{
	"mapf": {
		Name:     "mapf",
		Callback: command_mapf,
	},
	"mapb": {
		Name:     "mapb",
		Callback: command_mapb,
	},
	"help": {
		Name:     "help",
		Callback: command_help,
	},
	"exit": {
		Name:     "exit",
		Callback: command_exit,
	},
}

func command_mapf(cfg *config.Config) error {

	var locations pokeapi.Locations
	if cfg.Locations.Next == "" {
		locations = pokeapi.GetLocations(pokeapi.LocationAreasURI)
	} else {
		locations = pokeapi.GetLocations(cfg.Locations.Next)
	}

	cfg.Locations.Next = locations.Next
	cfg.Locations.Prev = locations.Previous
	for _, v := range locations.Results {
		fmt.Printf("%v \n", v.Name)
	}

	return nil
}

func command_mapb(cfg *config.Config) error {

	var locations pokeapi.Locations
	if cfg.Locations.Prev == "" {
		locations = pokeapi.GetLocations(pokeapi.LocationAreasURI)
	} else {
		locations = pokeapi.GetLocations(cfg.Locations.Prev)
	}

	cfg.Locations.Next = locations.Next
	cfg.Locations.Prev = locations.Previous
	for _, v := range locations.Results {
		fmt.Printf("%v \n", v.Name)
	}

	return nil
}

func command_help(cfg *config.Config) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	fmt.Println("- help: Displays a help message")
	fmt.Println("- mapf: Displays the next 20 locations")
	fmt.Println("- mapb: Displays the previous 20 locations")
	fmt.Println("- exit: Exit the Pokedex")

	return nil
}

func command_exit(cfg *config.Config) error {

	os.Exit(0)

	return nil
}

func GetCommand(name string) (*Command, error) {
	c, ok := commands[name]
	if !ok {
		return nil, errors.New("ERROR: There is no such command. To know how the application works type help")
	}

	return &c, nil
}
