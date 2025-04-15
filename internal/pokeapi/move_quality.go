package pokeapi

import (
	"math/rand"
	"sort"
	
	"github.com/FT1006/pokedexcli/internal/models"
)

// CalculateMoveQuality computes a quality score for a move based on its attributes
// This gives us a way to rank moves and select better ones for Pokemon
func CalculateMoveQuality(move MoveDetails) float64 {
	// Base quality starts with power, the most important factor
	quality := float64(move.Power)

	// Calculate accuracy factor (treats null accuracy as 100%)
	var accuracyFactor float64 = 1.0
	if move.Accuracy > 0 {
		accuracyFactor = float64(move.Accuracy) / 100.0
	}

	// Apply accuracy to the quality score
	quality *= accuracyFactor

	// Add bonus for priority moves (moves that go first)
	if move.Priority > 0 {
		quality += float64(move.Priority * 10)
	}

	// Add bonus for moves with special effects
	if move.EffectChance > 0 {
		effectBonus := float64(move.EffectChance) * 0.2
		quality += effectBonus
	}

	// For status moves (typically power=0), assign a moderate base quality
	if move.Power == 0 && move.DamageClass.Name == "status" {
		// Base quality for status moves
		quality = 50.0

		// Status moves with effects get a bonus
		if move.EffectChance > 0 {
			quality += float64(move.EffectChance) * 0.5
		}

		// Status moves with priority get a bigger bonus
		if move.Priority > 0 {
			quality += float64(move.Priority * 15)
		}
	}

	// Make sure we never return a negative quality
	if quality < 0 {
		quality = 0
	}

	return quality
}

// MoveWithQuality pairs a move with its quality score for easier sorting
type MoveWithQuality struct {
	Move   MoveDetails
	Quality float64
}

// SortMovesByQuality sorts moves by their quality score in descending order
func SortMovesByQuality(moves []MoveDetails) []MoveDetails {
	// Create a slice of moves with their quality scores
	movesWithQuality := make([]MoveWithQuality, len(moves))
	for i, move := range moves {
		movesWithQuality[i] = MoveWithQuality{
			Move:   move,
			Quality: CalculateMoveQuality(move),
		}
	}

	// Sort by quality (highest first)
	sort.Slice(movesWithQuality, func(i, j int) bool {
		return movesWithQuality[i].Quality > movesWithQuality[j].Quality
	})

	// Extract just the sorted moves
	sortedMoves := make([]MoveDetails, len(moves))
	for i, m := range movesWithQuality {
		sortedMoves[i] = m.Move
	}

	return sortedMoves
}

// SelectWeightedRandomMove selects a move using weighted randomness based on quality
// Higher quality moves have a higher chance of being selected
func SelectWeightedRandomMove(moves []MoveDetails) (MoveDetails, bool) {
	if len(moves) == 0 {
		return MoveDetails{}, false
	}

	if len(moves) == 1 {
		return moves[0], true
	}

	// Calculate quality scores and total quality
	movesWithQuality := make([]MoveWithQuality, len(moves))
	totalQuality := 0.0
	
	for i, move := range moves {
		quality := CalculateMoveQuality(move)
		movesWithQuality[i] = MoveWithQuality{Move: move, Quality: quality}
		totalQuality += quality
	}

	// Handle the case where total quality is 0 (all moves have 0 quality)
	if totalQuality == 0 {
		// Fall back to uniform random selection
		index := rand.Intn(len(moves))
		return moves[index], true
	}

	// Generate a random point along the quality spectrum
	r := rand.Float64() * totalQuality
	
	// Find which move corresponds to this point
	cumulativeQuality := 0.0
	for _, m := range movesWithQuality {
		cumulativeQuality += m.Quality
		if r <= cumulativeQuality {
			return m.Move, true
		}
	}

	// Fallback (should never reach here, but just in case)
	return movesWithQuality[len(movesWithQuality)-1].Move, true
}

// SelectQualityBasicMove selects a basic move using quality-weighted randomness
func SelectQualityBasicMove(basicMoves []MoveDetails) (MoveDetails, bool) {
	return SelectWeightedRandomMove(basicMoves)
}

// SelectQualitySpecialMove selects a special move using quality-weighted randomness
func SelectQualitySpecialMove(specialMoves []MoveDetails) (MoveDetails, bool) {
	return SelectWeightedRandomMove(specialMoves)
}

// GetTopNMoves returns the top N moves by quality (or all moves if fewer than N)
func GetTopNMoves(moves []MoveDetails, n int) []MoveDetails {
	sorted := SortMovesByQuality(moves)
	if len(sorted) <= n {
		return sorted
	}
	return sorted[:n]
}

// SelectQualityMoves is an enhanced version of SelectMoves that uses quality scoring
// It replaces the previous random selection with quality-weighted selection
func SelectQualityMoves(allMoves []MoveDetails) (models.Skill, models.Skill, error) {
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
		// If we have some moves but not in both categories, try to be more lenient
		if len(allMoves) > 0 {
			// Just sort all moves by quality and pick the best regardless of category
			sortedMoves := SortMovesByQuality(allMoves)
			if len(sortedMoves) >= 2 {
				// Pick the best two moves
				basicMove := sortedMoves[0]
				specialMove := sortedMoves[1]
				return ConvertMoveToSkill(basicMove), ConvertMoveToSkill(specialMove), nil
			} else if len(sortedMoves) == 1 {
				// If only one move is available, use it as both basic and special (better than nothing)
				move := sortedMoves[0]
				return ConvertMoveToSkill(move), ConvertMoveToSkill(move), nil
			}
		}
		// Ultimate fallback if no valid moves at all
		return emptySkill, emptySkill, nil
	}

	// For Pokemon with many moves, focus on the top quality moves
	// (helps prevent low-quality moves being picked just because a Pokemon has a lot of them)
	const topMovesCount = 5
	if len(basicMoves) > topMovesCount {
		basicMoves = GetTopNMoves(basicMoves, topMovesCount)
	}
	if len(specialMoves) > topMovesCount {
		specialMoves = GetTopNMoves(specialMoves, topMovesCount)
	}

	// Select moves with quality-weighted randomness
	basicMove, _ := SelectQualityBasicMove(basicMoves)
	specialMove, _ := SelectQualitySpecialMove(specialMoves)

	// Convert to Skill models
	basicSkill := ConvertMoveToSkill(basicMove)
	specialSkill := ConvertMoveToSkill(specialMove)

	return basicSkill, specialSkill, nil
}