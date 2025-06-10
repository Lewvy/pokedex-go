package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func CommandCatch(c *Config) error {
	if len(c.Args) != 2 {
		return fmt.Errorf("Invalid arguments. want=2, got=%d", len(c.Args))
	}
	pokemon := c.Args[1]
	var baseExp int
	var err error
	ok, baseExp, err := c.PokeapiClient.ValidatePokemon(pokemon)
	if !ok {
		return fmt.Errorf("Invalid pokemon\n")
	}
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	if ok := c.PokeapiClient.AttemptCatch(int(math.Abs(float64(baseExp)))); ok {
		fmt.Printf("%s was caught!\n", pokemon)
		poke := c.Pokemons[pokemon]
		poke.BaseExperience = baseExp
		c.Pokemons[pokemon] = poke
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	return nil
}

func CommandExplore(c *Config) error {
	if len(c.Args) != 2 {
		return fmt.Errorf("Invalid arguments. want=2, got=%d", len(c.Args))
	}
	city := c.Args[1]
	url := "https://pokeapi.co/api/v2/location-area/" + city
	fmt.Println(url)
	pokemonResp, err := c.PokeapiClient.ListPokemons(&url)
	for _, encounters := range pokemonResp.PokemonEncounters {
		fmt.Println(encounters.Pokemon.Name)
	}
	if err != nil {
		return err
	}
	return nil
}

func CommandMapb(c *Config) error {
	if c.PrevLocationURL == nil {
		return fmt.Errorf("You are on the first page")
	}
	url := *c.PrevLocationURL
	locations, err := c.PokeapiClient.ListLocations(&url)
	if err != nil {
		return err
	}

	c.NextLocationURL = locations.Next
	c.PrevLocationURL = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)

	}

	return nil
}

func CommandMap(c *Config) error {
	locations, err := c.PokeapiClient.ListLocations(c.NextLocationURL)
	if err != nil {
		return err
	}

	c.NextLocationURL = locations.Next
	c.PrevLocationURL = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func CommandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	m := GetCommands()
	for _, v := range m {
		fmt.Printf("%s: %s\n", v.Name, v.Description)
	}
	return nil
}

func CommandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func ClearScreen(cfg *Config) error {
	cmd := exec.Command("clear")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Print(string(out))
	return nil
}

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
