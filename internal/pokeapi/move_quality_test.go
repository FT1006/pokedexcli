package pokeapi

import (
	"testing"
)

func TestCalculateMoveQuality(t *testing.T) {
	// Test standard attack move
	tackle := MoveDetails{
		Name:     "tackle",
		Power:    40,
		Accuracy: 100,
		DamageClass: struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			Name: "physical",
		},
	}
	tackleQuality := CalculateMoveQuality(tackle)
	
	// 40 power * 1.0 accuracy = 40 quality
	if tackleQuality != 40 {
		t.Errorf("Expected tackle quality to be 40, got %f", tackleQuality)
	}

	// Test move with low accuracy
	lowAccuracyMove := MoveDetails{
		Name:     "dynamic-punch",
		Power:    100,
		Accuracy: 50,
		DamageClass: struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			Name: "physical",
		},
	}
	lowAccuracyQuality := CalculateMoveQuality(lowAccuracyMove)
	
	// 100 power * 0.5 accuracy = 50 quality
	if lowAccuracyQuality != 50 {
		t.Errorf("Expected low accuracy move quality to be 50, got %f", lowAccuracyQuality)
	}

	// Test status move
	statusMove := MoveDetails{
		Name:     "growl",
		Power:    0,
		Accuracy: 100,
		DamageClass: struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			Name: "status",
		},
	}
	statusQuality := CalculateMoveQuality(statusMove)
	
	// Base quality for status moves should be 50
	if statusQuality != 50 {
		t.Errorf("Expected status move quality to be 50, got %f", statusQuality)
	}

	// Test priority move
	priorityMove := MoveDetails{
		Name:     "quick-attack",
		Power:    40,
		Accuracy: 100,
		Priority: 1,
		DamageClass: struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			Name: "physical",
		},
	}
	priorityQuality := CalculateMoveQuality(priorityMove)
	
	// 40 power * 1.0 accuracy + 10 priority bonus = 50 quality
	if priorityQuality != 50 {
		t.Errorf("Expected priority move quality to be 50, got %f", priorityQuality)
	}

	// Test move with effect chance
	effectMove := MoveDetails{
		Name:        "flamethrower",
		Power:       90,
		Accuracy:    100,
		EffectChance: 10,
		DamageClass: struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			Name: "special",
		},
	}
	effectQuality := CalculateMoveQuality(effectMove)
	
	// 90 power * 1.0 accuracy + (10 * 0.2) effect bonus = 92 quality
	if effectQuality != 92 {
		t.Errorf("Expected effect move quality to be 92, got %f", effectQuality)
	}
}

func TestSortMovesByQuality(t *testing.T) {
	// Create a list of moves with varying qualities
	moves := []MoveDetails{
		{
			Name:     "tackle",
			Power:    40,
			Accuracy: 100,
		},
		{
			Name:     "hyper-beam",
			Power:    150,
			Accuracy: 90,
		},
		{
			Name:     "quick-attack",
			Power:    40,
			Accuracy: 100,
			Priority: 1,
		},
	}

	// Sort the moves
	sortedMoves := SortMovesByQuality(moves)

	// Verify the order
	if len(sortedMoves) != 3 {
		t.Errorf("Expected 3 sorted moves, got %d", len(sortedMoves))
	}

	// hyper-beam should be first (135 quality)
	if sortedMoves[0].Name != "hyper-beam" {
		t.Errorf("Expected hyper-beam to be first, got %s", sortedMoves[0].Name)
	}

	// quick-attack should be second (50 quality)
	if sortedMoves[1].Name != "quick-attack" {
		t.Errorf("Expected quick-attack to be second, got %s", sortedMoves[1].Name)
	}

	// tackle should be third (40 quality)
	if sortedMoves[2].Name != "tackle" {
		t.Errorf("Expected tackle to be third, got %s", sortedMoves[2].Name)
	}
}

func TestSelectWeightedRandomMove(t *testing.T) {
	// Test with empty list
	_, ok := SelectWeightedRandomMove([]MoveDetails{})
	if ok {
		t.Error("Expected SelectWeightedRandomMove to return false with empty list")
	}

	// Test with single move
	singleMove := []MoveDetails{
		{Name: "tackle", Power: 40, Accuracy: 100},
	}
	move, ok := SelectWeightedRandomMove(singleMove)
	if !ok || move.Name != "tackle" {
		t.Errorf("Expected to get 'tackle' from single move list, got %s", move.Name)
	}

	// Test with multiple moves
	// We can't test the randomness directly, but we can verify it doesn't crash
	multiMoves := []MoveDetails{
		{Name: "tackle", Power: 40, Accuracy: 100},
		{Name: "hyper-beam", Power: 150, Accuracy: 90},
		{Name: "quick-attack", Power: 40, Accuracy: 100, Priority: 1},
	}
	_, ok = SelectWeightedRandomMove(multiMoves)
	if !ok {
		t.Error("Expected SelectWeightedRandomMove to return true with multi-move list")
	}
}

func TestGetTopNMoves(t *testing.T) {
	// Create a list of moves
	moves := []MoveDetails{
		{Name: "tackle", Power: 40, Accuracy: 100},
		{Name: "hyper-beam", Power: 150, Accuracy: 90},
		{Name: "quick-attack", Power: 40, Accuracy: 100, Priority: 1},
		{Name: "flamethrower", Power: 90, Accuracy: 100},
		{Name: "thunder", Power: 110, Accuracy: 70},
	}

	// Get top 3 moves
	topMoves := GetTopNMoves(moves, 3)

	// Verify count
	if len(topMoves) != 3 {
		t.Errorf("Expected 3 top moves, got %d", len(topMoves))
	}

	// Verify they are ordered by quality
	if topMoves[0].Name != "hyper-beam" {
		t.Errorf("Expected first top move to be hyper-beam, got %s", topMoves[0].Name)
	}

	// When N > list size, should return all moves
	allMoves := GetTopNMoves(moves, 10)
	if len(allMoves) != 5 {
		t.Errorf("Expected GetTopNMoves to return all 5 moves when N=10, got %d", len(allMoves))
	}
}

func TestSelectQualityMoves(t *testing.T) {
	// Create test moves with clear quality differences
	testMoves := []MoveDetails{
		// Basic moves
		{
			Name:     "tackle",
			Power:    40,
			Accuracy: 100,
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "physical",
			},
			Type: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "normal",
			},
		},
		{
			Name:     "mega-punch",
			Power:    80,
			Accuracy: 85,
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "physical",
			},
			Type: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "normal",
			},
		},
		// Special moves
		{
			Name:     "ember",
			Power:    40,
			Accuracy: 100,
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "special",
			},
			Type: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "fire",
			},
		},
		{
			Name:     "flamethrower",
			Power:    90,
			Accuracy: 100,
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "special",
			},
			Type: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "fire",
			},
		},
	}

	// Run the selection function
	basicSkill, specialSkill, err := SelectQualityMoves(testMoves)
	if err != nil {
		t.Errorf("Unexpected error in SelectQualityMoves: %v", err)
	}

	// In a deterministic world, we would expect mega-punch (quality ~68) 
	// and flamethrower (quality 90) to be selected
	// But since there's randomness involved, we can only check for reasonable outputs
	
	// Basic skill should be either tackle or mega-punch
	if basicSkill.Name != "tackle" && basicSkill.Name != "mega-punch" {
		t.Errorf("Expected basic skill to be tackle or mega-punch, got %s", basicSkill.Name)
	}
	
	// Special skill should be either ember or flamethrower
	if specialSkill.Name != "ember" && specialSkill.Name != "flamethrower" {
		t.Errorf("Expected special skill to be ember or flamethrower, got %s", specialSkill.Name)
	}
	
	// Check for class correctness
	if basicSkill.Class != "physical" {
		t.Errorf("Expected basic skill class to be physical, got %s", basicSkill.Class)
	}
	if specialSkill.Class != "special" {
		t.Errorf("Expected special skill class to be special, got %s", specialSkill.Class)
	}
}

func TestSelectQualityMovesWithLimitedMoves(t *testing.T) {
	// Test with only special moves available
	onlySpecialMoves := []MoveDetails{
		{
			Name:     "flamethrower",
			Power:    90,
			Accuracy: 100,
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "special",
			},
			Type: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "fire",
			},
		},
		{
			Name:     "ember",
			Power:    40,
			Accuracy: 100,
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "special",
			},
			Type: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "fire",
			},
		},
	}
	
	// Should get best moves regardless of category if categories aren't balanced
	basicSkill, specialSkill, err := SelectQualityMoves(onlySpecialMoves)
	if err != nil {
		t.Errorf("Unexpected error in SelectQualityMoves with limited moves: %v", err)
	}
	
	// First skill should be flamethrower (higher quality)
	if basicSkill.Name != "flamethrower" {
		t.Errorf("Expected first skill to be flamethrower, got %s", basicSkill.Name)
	}
	
	// Second skill should be ember
	if specialSkill.Name != "ember" {
		t.Errorf("Expected second skill to be ember, got %s", specialSkill.Name)
	}
	
	// Test with only one move available
	singleMove := []MoveDetails{
		{
			Name:     "tackle",
			Power:    40,
			Accuracy: 100,
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "physical",
			},
			Type: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "normal",
			},
		},
	}
	
	// Should use the same move for both slots
	basicSkill, specialSkill, err = SelectQualityMoves(singleMove)
	if err != nil {
		t.Errorf("Unexpected error in SelectQualityMoves with single move: %v", err)
	}
	
	if basicSkill.Name != "tackle" || specialSkill.Name != "tackle" {
		t.Errorf("Expected both skills to be tackle, got %s and %s", 
			basicSkill.Name, specialSkill.Name)
	}
	
	// Test with no moves at all
	noMoves := []MoveDetails{}
	basicSkill, specialSkill, err = SelectQualityMoves(noMoves)
	if err != nil {
		t.Errorf("Unexpected error in SelectQualityMoves with no moves: %v", err)
	}
	
	// Should fall back to struggle
	if basicSkill.Name != "struggle" || specialSkill.Name != "struggle" {
		t.Errorf("Expected both skills to be struggle, got %s and %s", 
			basicSkill.Name, specialSkill.Name)
	}
}