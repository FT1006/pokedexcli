package pokeapi

import (
	"github.com/FT1006/pokedexcli/internal/models"
)

// MoveCategorization contains methods for categorizing Pokemon moves
// into basic and special categories based on their properties

// CategorizeMoves separates move details into basic and special categories
// Basic moves are typically physical attacks with lower power
// Special moves are typically special attacks or high-power attacks
func CategorizeMoves(moves []MoveDetails) ([]MoveDetails, []MoveDetails) {
	var basicMoves []MoveDetails
	var specialMoves []MoveDetails

	for _, move := range moves {
		// Skip moves with no power (status moves)
		if move.Power <= 0 {
			continue
		}

		// Basic categorization logic
		if move.DamageClass.Name == "physical" {
			basicMoves = append(basicMoves, move)
		} else if move.DamageClass.Name == "special" {
			specialMoves = append(specialMoves, move)
		}
		// Status moves generally don't deal damage, so we exclude them
	}

	return basicMoves, specialMoves
}

// ConvertMoveToSkill converts a MoveDetails object to a Skill model
func ConvertMoveToSkill(move MoveDetails) models.Skill {
	return models.Skill{
		Name:   move.Name,
		URL:    getBaseURL() + "move/" + move.Name,
		Damage: move.Power,
		Type:   move.Type.Name,
		Class:  move.DamageClass.Name,
	}
}

// HasValidMoves checks if the categorized moves contain at least one valid move
// in each category. If not, it suggests whether to relax categorization criteria.
func HasValidMoves(basicMoves, specialMoves []MoveDetails) bool {
	return len(basicMoves) > 0 && len(specialMoves) > 0
}

// GetMoveCounts returns the count of moves in each category for debugging or reporting
func GetMoveCounts(basicMoves, specialMoves []MoveDetails) (int, int) {
	return len(basicMoves), len(specialMoves)
}