package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := getCommands()[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
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
	callback			func(*config, ...string) 	error
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
			description: "Lists next locations",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Lists previous locations",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Get a list of all the pokemon in the area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch a pokemon",
			callback: commandCatch,
		},
	}
}