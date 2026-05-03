package main

type Config struct {
	PokeAPI  string
	Next     string
	Previous string
}

var config = &Config{
	PokeAPI: "https://pokeapi.co/api/v2/",
}
