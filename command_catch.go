package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(config *config, pokemonName *string) error {
	if pokemonName == nil {
		return fmt.Errorf(`Include the name of the Pokemon you wish to catch, example: "catch pikachu"`)
	}

	pokemon, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fixedName := strings.ToUpper(pokemon.Name[:1]) + pokemon.Name[1:]

	fmt.Printf("Throwing a pokeball at %s...\n", fixedName)

	if (rand.Intn(500) > pokemon.BaseExperience) && (rand.Intn(500) > pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", fixedName)
		config.pokedex[*pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", fixedName)
	}

	return nil
}
