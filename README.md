# Pokédex CLI

A command-line Pokédex built in Go. Explore the Pokémon world, catch Pokémon, and inspect your collection — all from your terminal.

Data is fetched live from the [PokéAPI](https://pokeapi.co/).

## Features

- Browse Pokémon location areas with pagination
- Explore a specific area to see what Pokémon live there
- Catch Pokémon with a probability-based catch system
- Inspect caught Pokémon to view their stats, types, height and weight
- View your full Pokédex collection
- API response caching to avoid redundant network requests

## Commands

| Command | Description |
|---|---|
| `help` | Show all available commands |
| `map` | Display the next 20 location areas |
| `mapb` | Display the previous 20 location areas |
| `explore <area>` | List all Pokémon found in a given area |
| `catch <pokemon>` | Attempt to catch a Pokémon |
| `inspect <pokemon>` | View stats, types, height and weight of a caught Pokémon |
| `pokedex` | List all Pokémon you have caught |
| `exit` | Exit the program |

## How catching works

Each Pokémon has a `base_experience` value from the PokéAPI that reflects its overall strength. The catch probability is calculated as:

```
catchRate = 1.0 - baseExperience / 400.0
```

Weaker Pokémon (low base experience) are easier to catch. Stronger Pokémon are harder. No Pokémon is impossible to catch — just unlikely.

## Getting started

**Requirements:** Go 1.21+

```bash
git clone https://github.com/lillysilly3/pokedex
cd pokedex
go run .
```

## Example session

```
Pokédex > map
cerulean-city
vermilion-city
lavender-town
...

Pokédex > explore cerulean-city
Exploring cerulean-city...
Found Pokemon:
 - psyduck
 - golduck
 - poliwag

Pokédex > catch psyduck
Throwing a Pokeball at psyduck...
psyduck was caught!
You may now inspect it with the inspect command.

Pokédex > inspect psyduck
Name: psyduck
Height: 8
Weight: 196
Stats:
 -hp: 50
 -attack: 52
 -defense: 48
Types:
 -water
```

## Project structure

```
.
├── main.go          # Entry point, sets up the REPL and API client
├── repl.go          # Read-Eval-Print Loop logic
├── commands.go      # All CLI command implementations
├── repl_test.go     # Tests for the REPL
└── internal/
    └── pokeapi/     # PokéAPI client with caching
```
