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

![Pokedex screenshot 1](/internal/images/pokedex1.png)
![Pokedex screenshot 2](/internal/images/pokedex2.png)

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

## What I Learned

- **Go basics** - Structs, interfaces, error handling, and Go's type system coming from Python
- **REST API integration** - Fetching and deserializing JSON responses from PokeAPI
- **Caching** - Implementing a time-based cache to avoid redundant API calls
- **Concurrency** - Using goroutines to handle background tasks
- **REPL design** - Building an interactive command loop with a clean command dispatch pattern
- **CLI UX** - Designing intuitive commands and pagination for navigating large datasets
- **Testing in Go** - Writing unit tests with Go's built-in testing package

## Getting started

**Requirements:** Go 1.21+

```bash
git clone https://github.com/lillysilly3/pokedex
cd pokedex
go run .
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

## Future ideas

```
 - Update the CLI to support the "up" arrow to cycle through previous commands
 - Refactor code to organize it better and make it more testable
 - Keep pokemon in a "party" and allow them to level up
 - Persist a user's Pokedex to disk so they can save progress between sessions
 - Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
```

## Acknowledgments

This project was built as part of the [Boot.dev](https://boot.dev) curriculum.
