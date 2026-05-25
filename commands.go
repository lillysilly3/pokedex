package main

import (
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"errors"
	"math/rand"
)

type Config struct {
	Next		*string
	Previous	*string
	Client		*pokeapi.Client
	Pokedex		map[string]pokeapi.Pokemon
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
		"catch": {
		name:			"catch",
		description:	"Attemps to catch a pokemon and add it to your Pokedex",
		callback:		commandCatch,
		},
		"inspect": {
		name:			"inspect",
		description:	"Takes the name of a Pokemon and prints the name, height, weight, stats and type(s) of the Pokemon",
		callback:		commandInspect,
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
		return errors.New("Wrong command usage")
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

func commandCatch(cfg *Config, args []string) error {
	if len(args) != 1 {
		return errors.New("Wrong command usage")
	}
	pokemonName := args[0]

	pokemon, err := cfg.Client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	
	catchRate := 1.0 - float64(pokemon.BaseExperience)/400.0
	randVal := rand.Float64()
	
	if randVal <= catchRate {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}


	return nil
}

func commandInspect(cfg *Config, args []string) error {
	if len(args) != 1 {
		return errors.New("Wrong command usage")
	}
	pokemonName := args[0]
	
	pokemon, ok := cfg.Pokedex[pokemonName]
	if !ok {
		fmt.Printf("you have not caught %s\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n",pokemon.Height)
	fmt.Printf("Weight: %d\n",pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeP := range pokemon.Types {
		fmt.Printf(" -%s\n", typeP.Type.Name)
	}
	
	return nil
}