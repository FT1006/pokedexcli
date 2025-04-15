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

// GetMoveDetailsBatch fetches details for multiple moves in parallel
// This optimizes API calls when we need to fetch several moves at once
func (c *Client) GetMoveDetailsBatch(moveNames []string) (map[string]MoveDetails, error) {
	results := make(map[string]MoveDetails)
	errChan := make(chan error, len(moveNames))
	resultChan := make(chan struct {
		name string
		move MoveDetails
	}, len(moveNames))
	// Define the worker function that fetches a single move
	worker := func(name string) {
		move, err := c.GetMoveDetails(name)
		if err != nil {
			errChan <- fmt.Errorf("error fetching move %s: %w", name, err)
			return
		}
		resultChan <- struct {
			name string
			move MoveDetails
		}{name, move}
	}
	// Launch workers for each move
	for _, name := range moveNames {
		go worker(name)
	}
	// Collect results
	for i := 0; i < len(moveNames); i++ {
		select {
		case result := <-resultChan:
			results[result.name] = result.move
		case err := <-errChan:
			return nil, err
		}
	}
	return results, nil
}

// PrefetchCommonMoves preloads common moves into the cache to improve responsiveness
// This is meant to be called in a background goroutine during app initialization
func (c *Client) PrefetchCommonMoves() {
	// List of common moves to prefetch
	commonMoves := []string{
		"tackle", "scratch", "pound", "quick-attack", "slam", // Common physical moves
		"thunderbolt", "flamethrower", "water-gun", "ice-beam", "solar-beam", // Common special moves
		"hyper-beam", "earthquake", "psychic", "surf", "fire-blast", // Strong moves
	}
	fmt.Println("Prefetching common moves in background...")
	// Use a separate goroutine to avoid blocking
	go func() {
		_, err := c.GetMoveDetailsBatch(commonMoves)
		if err != nil {
			fmt.Println("Error prefetching moves:", err)
		} else {
			fmt.Println("Successfully prefetched", len(commonMoves), "common moves")
		}
	}()
}
