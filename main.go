package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		
		if len(input) == 0 {
			continue
		}
		
		commandName := input[0]
		command, exists := getCommands()[commandName]	
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}