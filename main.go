package main

import (
	"time"

	"github.com/pshunk/pokedex/internal/pokecache/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Minute * 5)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
