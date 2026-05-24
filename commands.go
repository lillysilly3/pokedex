package main

import (
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"errors"
)

type Config struct {
	Next		*string
	Previous	*string
	Client		*pokeapi.Client
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
		name:			"exit",
		description:	"Exit the Pokedex",
		callback:		commandExit,
		},
		"help": {
		name:			"help",
		description:	"Displays a help message",
		callback:		commandHelp,
		},
		"map": {
		name:			"map",
		description:	"Displays the names of next 20 location areas",
		callback:		commandMap,
		},
		"mapb": {
		name:			"mapb",
		description:	"Displays the names of previous 20 location areas",
		callback:		commandMapb,
		},
		"explore": {
		name:			"explore",
		description: 	"Displays the names of pokemon in given area",
		callback:		commandExplore,
		},
	}
}

func commandExit(_ *Config, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *Config, _ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap(cfg *Config, _ []string) error {
	locationResp, err := cfg.Client.ListLocations(cfg.Next)
	if err != nil {
		return err
	}
	
	for _, r := range locationResp.Results {
		fmt.Println(r.Name)
	}
	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous
		
	return nil
}

func commandMapb(cfg *Config, _ []string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	
	url := *cfg.Previous
	locationResp, err := cfg.Client.ListLocations(&url)
	if err != nil {
		return err
	}

	for _, r := range locationResp.Results {
		fmt.Println(r.Name)
	}
	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous
		
	return nil
}

func commandExplore(cfg *Config, args []string) error {
	if len(args) != 1 {
		return errors.New("Wrogn command usage")
	}
	areaName := args[0]
	
	pokemonResp, err := cfg.Client.GetLocationArea(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")

	for _, encounter := range pokemonResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	} 
	return nil
}