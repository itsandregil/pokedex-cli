package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"os"

	"github.com/itsandregil/pokedex-cli/internal/pokeapi"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *Config, args ...string) error {
	locationArea, err := cfg.pokeAPIClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationArea.Next
	cfg.prevLocationURL = locationArea.Previous

	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(cfg *Config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you are at the first page")
	}

	locationArea, err := cfg.pokeAPIClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationArea.Next
	cfg.prevLocationURL = locationArea.Previous

	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: explore <location-name>")
	}
	locationName := args[0]

	result, err := cfg.pokeAPIClient.ListEncounters(locationName)
	if err != nil {
		return fmt.Errorf("error exploring area: %v", err)
	}

	for _, encounter := range result.Encounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon-name>")
	}
	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := cfg.pokeAPIClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("error catching pokemon: %v", err)
	}

	// Using the exponential decay formula
	catchChance := 0.05 + (0.9-0.05)*math.Exp(-0.01*float64(pokemon.BaseExperience))
	userRoll := rand.Float64()

	if userRoll <= catchChance {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemon
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: inspect <pokemon-name>")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("You have not caught %s yet!", pokemonName)
	}
	printPokemon(pokemon)
	return nil
}

func printPokemon(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}
}

func commandPokedex(cfg *Config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your pokedex is empty!")
		return nil
	}

	for pokemonName := range cfg.pokedex {
		fmt.Printf("- %s\n", pokemonName)
	}
	return nil
}
