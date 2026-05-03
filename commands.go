package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
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

func commandMap(cfg *Config) error {
	url := cfg.PokeAPI + "location-area"
	if cfg.Next != "" {
		url = cfg.Next
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locationArea LocationArea
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationArea); err != nil {
		return err
	}

	if locationArea.Next != nil {
		cfg.Next = *locationArea.Next
	}
	if locationArea.Previous != nil {
		cfg.Previous = *locationArea.Previous
	}
	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("You are at the first page")
		return nil
	}

	res, err := http.Get(cfg.Previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locationArea LocationArea
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationArea); err != nil {
		return err
	}

	if locationArea.Next != nil {
		cfg.Next = *locationArea.Next
	}
	if locationArea.Previous != nil {
		cfg.Previous = *locationArea.Previous
	}
	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}
