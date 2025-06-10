package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	Client := pokeapi.NewClient(time.Second*5, time.Minute*5)
	PokemonMap := make(map[string]Pokemon)

	c := Config{
		PokeapiClient: Client,
		Pokemons:      PokemonMap,
	}
	Start(&c)

}
