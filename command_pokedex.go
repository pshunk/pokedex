package main

import (
	"fmt"
	"strings"
)

func commandPokedex(config *config, param *string) error {
	if len(config.pokedex) == 0 {
		return fmt.Errorf("Your pokedex is empty, try catching some Pokemon")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.pokedex {
		fmt.Printf("  - %s\n", strings.ToUpper(pokemon.Name[:1])+pokemon.Name[1:])
	}
	return nil
}
