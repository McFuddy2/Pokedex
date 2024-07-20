package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// Check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// Found the data in the cache!
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil{
			return Pokemon{}, err
		}
		return pokemon, nil
	}


	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil{
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil{
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil{
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
