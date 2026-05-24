package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

func (c *Client) GetLocationArea(areaName string) (PokemonInArea, error) {
	url := baseURL + "/location-area/" + areaName
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := PokemonInArea{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return PokemonInArea{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInArea{}, err
	}
	defer resp.Body.Close()

	//Checking if resp actually legit
	if resp.StatusCode > 299 {
    	return PokemonInArea{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInArea{}, err
	}
	
	pokemonResp := PokemonInArea{}
	
	err = json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return PokemonInArea{}, err
	}
	c.cache.Add(url, body)

	return pokemonResp, nil
}