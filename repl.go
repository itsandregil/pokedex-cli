package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exist the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays locations to explore in an area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays locations to explore in the previous area",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore pokemons in a given location area",
			callback:    commandExplore,
		},
	}
}

func startREPL(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		commandArgs := words[1:]
		command, exits := getCommands()[commandName]
		if exits {
			err := command.callback(cfg, commandArgs...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
