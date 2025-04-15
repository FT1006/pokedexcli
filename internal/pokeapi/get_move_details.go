package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetMoveDetails fetches detailed information about a specific move from the PokeAPI
func (c *Client) GetMoveDetails(moveName string) (MoveDetails, error) {
	url := getBaseURL() + "move/" + moveName
	
	// Check the cache first
	if cachedData, ok := c.pokecache.Get(url); ok {
		fmt.Println("Cache hit for move:", moveName)
		var moveDetails MoveDetails
		err := json.Unmarshal(cachedData, &moveDetails)
		if err != nil {
			return MoveDetails{}, err
		}
		return moveDetails, nil
	}
	
	// Cache miss, make the API request
	fmt.Println("Getting move details for:", moveName)
	res, err := http.Get(url)
	if err != nil {
		return MoveDetails{}, err
	}
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return MoveDetails{}, err
	}
	
	// Store the raw response in the cache
	c.pokecache.Add(url, body)
	
	// Parse the response
	var moveDetails MoveDetails
	err = json.Unmarshal(body, &moveDetails)
	if err != nil {
		return MoveDetails{}, err
	}
	
	return moveDetails, nil
}

// GetPokemonMoves fetches all moves a Pokemon can learn from the PokeAPI
func (c *Client) GetPokemonMoves(pokemonName string) ([]PokemonMove, error) {
	// Leverage the existing GetPokemonDetails method to get Pokemon data
	pokemonDetails, err := c.GetPokemonDetails(pokemonName)
	if err != nil {
		return nil, fmt.Errorf("failed to get Pokemon moves: %w", err)
	}
	
	return pokemonDetails.Moves, nil
}