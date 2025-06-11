package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
)

const PROMPT = "Pokedex > "

type Config struct {
	PokeapiClient   pokeapi.Client
	NextLocationURL *string
	PrevLocationURL *string
	Args            []string
	Pokemons        map[string]*pokeapi.Pokemon
}

func Start(c *Config) {
	sc := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print(PROMPT)
		sc.Scan()
		c.Args = CleanInput(sc.Text())
		m := GetCommands()
		if command, ok := m[c.Args[0]]; ok {
			err := command.Callback(c)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
		}
	}

}

func GetCommands() map[string]cliCommand {
	m := map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},

		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},

		"map": {
			Name:        "map",
			Description: "Displays the names of areas in the pokemon world",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous location-areas",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the pokemon found in the provided area - usage: explore <location-area>",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempts to catch a pokemon - usage: catch <pokemon>",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspects caught pokemons",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Lists out all the acquired pokemons",
			Callback:    CommandPokedex,
		},
		"clear": {
			Name:        "clear",
			Description: "Clears the screen",
			Callback:    ClearScreen,
		},
	}
	return m
}
