package main

import (
	"fmt"
	"strings"
)

func commandInspect(config *config, pokemonName *string) error {
	if pokemonName == nil {
		return fmt.Errorf(`Include the name of the Pokemon you wish to view, example: "inspect pikachu"`)
	}

	pokemon, ok := config.pokedex[*pokemonName]
	if !ok {
		return fmt.Errorf("You don't have that pokemon in your pokedex, try catching it first")
	}
	fmt.Printf("Name: %s\n", strings.ToUpper(pokemon.Name[:1])+pokemon.Name[1:])
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
