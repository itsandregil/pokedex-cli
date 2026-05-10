package main

import (
	"errors"
	"fmt"
	"os"
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
