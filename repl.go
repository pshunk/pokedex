package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pshunk/pokedex/internal/pokecache/pokeapi"
)

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Err() != nil {
			continue
		}
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}
		command := userInput[0]
		if c, ok := config.commands[command]; ok {
			var err error
			if len(userInput) > 1 {
				err = c.callback(config, &userInput[1])
			} else {
				err = c.callback(config, nil)
			}
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	cleanedText := strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")
	return cleanedText
}

type command struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays locations in the Pokemon games, use the command again to advance the list",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Moves backwards in the list of map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays all of the Pokemon available in an area",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
}

type config struct {
	commands      map[string]command
	pokeapiClient pokeapi.Client
	mapUrl        string
	mapPrevUrl    string
}
