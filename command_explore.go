package main

import "fmt"

func commandExplore(config *config, locationName *string) error {
	if locationName == nil {
		return fmt.Errorf(`Include the name of the location you wish to explore, example: "explore canalave-city-area"`)
	}

	location, err := config.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *locationName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return err
}
