package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type RespShallowLocations struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
type RespPokemons struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListLocations(url *string) (RespShallowLocations, error) {
	pageURL := "https://pokeapi.co/api/v2/location-area"
	if url == nil {
		url = &pageURL
	}
	if val, ok := c.cache.Get(*url); ok {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err == nil {
			return locationResp, nil
		}
	}
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	locationResp := RespShallowLocations{}
	data, err := io.ReadAll(res.Body)
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	c.cache.Add(*url, data)

	return locationResp, nil
}

func (c *Client) ListPokemons(url *string) (RespPokemons, error) {
	pageUrl := *url

	pokemons := RespPokemons{}
	if val, ok := c.cache.Get(pageUrl); ok {
		err := json.Unmarshal(val, &pokemons)
		if err == nil {
			return pokemons, nil
		}
	}
	req, err := http.NewRequest("GET", pageUrl, nil)
	if err != nil {
		return RespPokemons{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemons{}, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemons{}, err
	}
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return RespPokemons{}, err
	}
	c.cache.Add(pageUrl, data)
	return pokemons, nil
}

func (c Client) AttemptCatch(baseExp int) bool {
	randNo := rand.Intn(baseExp)
	if randNo >= baseExp/3 {
		return true
	}
	return false
}

func (c Client) ValidatePokemon(pokemon string) (bool, int, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	if val, ok := c.cache.Get(url); ok {

		poke := map[string]any{}
		json.Unmarshal(val, &poke)
		x := int(poke["base_experience"].(float64))
		return true, x, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, 0, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return false, 0, err
	}
	if res.StatusCode >= 400 {
		return false, 0, fmt.Errorf("%s not found", pokemon)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return false, 0, err
	}
	p := struct {
		BaseExperience int `json:"base_experience"`
	}{}

	json.Unmarshal(data, &p)

	c.cache.Add(url, data)
	return true, p.BaseExperience, nil
}
