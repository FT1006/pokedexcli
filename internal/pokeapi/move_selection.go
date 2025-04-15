package pokeapi

import (
	"math/rand"

	"github.com/FT1006/pokedexcli/internal/models"
)

// Note: In Go 1.20+, rand is automatically seeded, so we no longer need 
// to seed it manually in an init function.

// SelectRandomBasicMove randomly selects a move from the basic moves category
// Returns the selected move and a boolean indicating success
func SelectRandomBasicMove(basicMoves []MoveDetails) (MoveDetails, bool) {
	if len(basicMoves) == 0 {
		return MoveDetails{}, false
	}
	
	index := rand.Intn(len(basicMoves))
	return basicMoves[index], true
}

// SelectRandomSpecialMove randomly selects a move from the special moves category
// Returns the selected move and a boolean indicating success
func SelectRandomSpecialMove(specialMoves []MoveDetails) (MoveDetails, bool) {
	if len(specialMoves) == 0 {
		return MoveDetails{}, false
	}
	
	index := rand.Intn(len(specialMoves))
	return specialMoves[index], true
}

// SelectMoves takes all of a Pokemon's moves, categorizes them, and selects
// one basic and one special move to assign to the Pokemon
func SelectMoves(allMoves []MoveDetails) (models.Skill, models.Skill, error) {
	// Default "empty" skills in case we can't find appropriate moves
	emptySkill := models.Skill{
		Name:   "struggle",
		URL:    getBaseURL() + "move/struggle",
		Damage: 50,
		Type:   "normal",
		Class:  "physical",
	}

	// Categorize the moves
	basicMoves, specialMoves := CategorizeMoves(allMoves)

	// Check if we have valid moves in both categories
	if !HasValidMoves(basicMoves, specialMoves) {
		// Handle Pokemon with limited movesets by being more lenient
		// For now, we'll just use the empty skill as a fallback
		return emptySkill, emptySkill, nil
	}

	// Select a random move from each category
	basicMove, _ := SelectRandomBasicMove(basicMoves)
	specialMove, _ := SelectRandomSpecialMove(specialMoves)

	// Convert to Skill models
	basicSkill := ConvertMoveToSkill(basicMove)
	specialSkill := ConvertMoveToSkill(specialMove)

	return basicSkill, specialSkill, nil
}

// GetMovesForPokemon is a convenience function that fetches a Pokemon's moves
// from the API and selects appropriate basic and special moves
func (c *Client) GetMovesForPokemon(pokemonName string) (models.Skill, models.Skill, error) {
	// Get Pokemon details to access its moves
	pokemonDetails, err := c.GetPokemonDetails(pokemonName)
	if err != nil {
		return models.Skill{}, models.Skill{}, err
	}

	// Extract move references from the Pokemon
	var moveRefs []string
	for _, move := range pokemonDetails.Moves {
		moveRefs = append(moveRefs, move.Move.Name)
	}

	// Fetch details for all moves
	moveDetailsMap, err := c.GetMoveDetailsBatch(moveRefs)
	if err != nil {
		return models.Skill{}, models.Skill{}, err
	}

	// Convert map to slice for processing
	var moveDetails []MoveDetails
	for _, details := range moveDetailsMap {
		moveDetails = append(moveDetails, details)
	}

	// Select the moves using quality-based selection
	return SelectQualityMoves(moveDetails)
}