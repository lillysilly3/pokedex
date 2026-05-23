package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	cfg := &Config{}
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		
		if len(input) == 0 {
			continue
		}
		
		commandName := input[0]
		command, ok := getCommands()[commandName]	
		if ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}