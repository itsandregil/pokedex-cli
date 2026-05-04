package main

import "github.com/itsandregil/pokedex-cli/internal/pokeapi"

type Config struct {
	pokeAPIClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}
