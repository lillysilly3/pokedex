package main

import (
    "strings"
	"fmt"
	"os"
	"net/http"
	"encoding/json"
)

type Config struct {
	Next *string
	Previous *string
}

type Result struct {
	Name	string		`json:"name"`
	Url		string		`json:"url"`
}

type LocationResponse struct {
	Count		int			`json:"count"`
	Next		*string		`json:"next"`
	Previous	*string		`json:"previous"`
	Results		[]Result	`json:"results"`
}

type cliCommand struct {
	name		string
	description	string
	callback	func(cfg *Config) error
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
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var location LocationResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&location); err != nil {
		return fmt.Errorf("getting location areas: %w", err)
	}
	

	for _, r := range location.Results {
		fmt.Println(r.Name)
	}
	cfg.Next = location.Next
	cfg.Previous = location.Previous
		
	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *cfg.Previous
	
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var location LocationResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&location); err != nil {
		return fmt.Errorf("getting location areas: %w", err)
	}
	

	for _, r := range location.Results {
		fmt.Println(r.Name)
	}
	cfg.Next = location.Next
	cfg.Previous = location.Previous
		
	return nil
}