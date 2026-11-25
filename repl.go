package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
			err := c.callback(config)
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
	callback    func(*config) error
}

func getCommands() map[string]command {
	return map[string]command{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

type config struct {
	commands map[string]command
}
