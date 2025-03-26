package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocationAreaList(url string) (LocationArea, error) {
	if url == "" {
		url = getBaseURL() + "location-area/"
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var areas LocationArea
	err = json.Unmarshal(body, &areas)
	if err != nil {
		log.Fatal(err)
	}
	return areas, nil
}
