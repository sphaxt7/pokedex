package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			if err := command.callback(); err != nil {
				fmt.Println(err)
			}
			continue

		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}


func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	splitText := strings.Fields(lowerText)
	
	return splitText
}


type cliCommand struct {
	name, description 	string
	callback			func() 	error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the map",
			callback: commandMap,
		},

	}
}