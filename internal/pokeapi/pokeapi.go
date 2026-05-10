package pokeapi

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

type LocationArea struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounters struct {
	Name       string   `json:"name"`
	Location   Location `json:"location"`
	Encounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
