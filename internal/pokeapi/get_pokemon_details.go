package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(name string) (PokemonRes, error) {
	var pokemonRes PokemonRes

	url := getBaseURL() + "pokemon/" + name

	// When making an API request:
	fmt.Println("Making API request to:", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Pokemon details request failed")
		return PokemonRes{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("failed to read body")
		return PokemonRes{}, err
	}

	err = json.Unmarshal(body, &pokemonRes)
	if err != nil {
		fmt.Println("failed to unmarshal")
		return PokemonRes{}, err
	}

	return pokemonRes, nil
}
