package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	Client := pokeapi.NewClient(time.Second*5, time.Minute*5)
	PokemonMap := make(map[string]*pokeapi.Pokemon)

	c := Config{
		PokeapiClient: Client,
		Pokemons:      PokemonMap,
	}
	Start(&c)

}
