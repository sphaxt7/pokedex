package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *Client) GetPokemon(PokemonName string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + PokemonName
	
	// caching
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("got a cache")
		MyPokemon := Pokemon{}
		if err := json.Unmarshal(data, &MyPokemon); err != nil {
			return Pokemon{}, err
		}
	
		return MyPokemon, nil		
	}
	fmt.Println("no cache")

	
	// requesting
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return Pokemon{}, fmt.Errorf("error: %v", res.StatusCode)
	}


	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	MyPokemon := Pokemon{}
	if err := json.Unmarshal(data, &MyPokemon); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return MyPokemon, nil
} 