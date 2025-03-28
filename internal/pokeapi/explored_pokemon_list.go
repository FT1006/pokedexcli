package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetExploredPokemonList(area string) (Explored, error) {
	var exploreds Explored

	url := getBaseURL() + "location-area/" + area

	if cached, ok := c.pokecache.Get(url); ok {
		// When retrieving from cache:
		fmt.Println("Cache hit! Using cached data for:", url)
		err := json.Unmarshal(cached, &exploreds)
		if err != nil {
			return Explored{}, err
		}
		return exploreds, nil
	}
	// When making an API request:
	fmt.Println("Cache miss! Making API request to:", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("area not found or failed to get the area")
		return Explored{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("failed to read body")
		return Explored{}, err
	}

	c.pokecache.Add(url, body)
	err = json.Unmarshal(body, &exploreds)
	if err != nil {
		fmt.Println("failed to unmarshal")
		return Explored{}, err
	}
	return exploreds, nil
}
