package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName *string) (Pokemon, error) {
	pokemon := Pokemon{}

	if pokemonName == nil {
		return pokemon, nil
	}

	url := baseURL + "/pokemon/" + *pokemonName

	if bytes, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(bytes, &pokemon)
		return pokemon, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemon, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon, err
	}
	if res.StatusCode > 299 {
		return pokemon, fmt.Errorf("Could not find the requested pokemon")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemon, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokemon, err
	}

	c.cache.Add(url, data)

	return pokemon, err
}
