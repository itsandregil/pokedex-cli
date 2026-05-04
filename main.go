package main

import (
	"github.com/itsandregil/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient()
	config := &Config{
		pokeAPIClient: pokeClient,
	}
	startREPL(config)
}
