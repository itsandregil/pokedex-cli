package main

import (
	"time"

	"github.com/itsandregil/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 15*time.Second)
	config := &Config{
		pokeAPIClient: pokeClient,
	}
	startREPL(config)
}
