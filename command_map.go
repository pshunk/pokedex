package main

import (
	"fmt"
)

func commandMapf(config *config, param *string) error {
	if config.mapUrl == "" {
		config.mapUrl = "https://pokeapi.co/api/v2/location-area/"
	}
	locations, err := config.pokeapiClient.GetLocations(&config.mapUrl)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	if locations.Next != nil {
		config.mapUrl = *locations.Next
	} else {
		config.mapUrl = ""
	}
	if locations.Previous != nil {
		config.mapPrevUrl = *locations.Previous
	} else {
		config.mapPrevUrl = ""
	}
	return nil
}

func commandMapb(config *config, param *string) error {
	if config.mapPrevUrl == "" {
		return fmt.Errorf("You're on the first page")
	}
	locations, err := config.pokeapiClient.GetLocations(&config.mapPrevUrl)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	if locations.Next != nil {
		config.mapUrl = *locations.Next
	} else {
		config.mapUrl = ""
	}
	if locations.Previous != nil {
		config.mapPrevUrl = *locations.Previous
	} else {
		config.mapPrevUrl = ""
	}
	return nil
}
