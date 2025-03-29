package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocationAreaList(url string) (LocationArea, error) {
	var areas LocationArea
	if url == "" {
		url = getBaseURL() + "location-area/"
	}

	if cached, ok := c.pokecache.Get(url); ok {
		// When retrieving from cache:
		fmt.Println("Cache hit! Using cached data for:", url)
		err := json.Unmarshal(cached, &areas)
		if err != nil {
			return LocationArea{}, err
		}
		return areas, nil
	}
	// When making an API request:
	fmt.Println("Cache miss! Making API request to:", url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	c.pokecache.Add(url, body)

	err = json.Unmarshal(body, &areas)
	if err != nil {
		log.Fatal(err)
	}
	return areas, nil
}
