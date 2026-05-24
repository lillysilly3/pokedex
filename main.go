package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(30*time.Second, 5*time.Minute)
	cfg := &Config{
		Client: &client,
	}
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		
		if len(input) == 0 {
			continue
		}
		
		commandName := input[0]
		args := input[1:]
		command, ok := getCommands()[commandName]	
		if ok {
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}