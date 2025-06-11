package main

import (
	"fmt"
)

func CommandPokedex(c *Config) error {
	fmt.Println("Your Pokedex:")
	for _, v := range c.Pokemons {
		fmt.Println("-", v.Species.Name)
	}
	return nil
}
