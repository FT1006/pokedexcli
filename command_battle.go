package main

import (
	"fmt"
)

func commandBattle(pokemonName string, c *Config) error {
	// Simply show battle information to the user
	showBattleInfo()
	return nil
}

// Show battle information to the user
func showBattleInfo() {
	fmt.Println("\n===== POKEMON BATTLE SYSTEM =====")
	fmt.Println("\nBATTLE WORKFLOW:")
	fmt.Println("1. A Pokemon is randomly chosen from your party")
	fmt.Println("2. Your Pokemon always attacks first")
	fmt.Println("3. You can choose between basic attack or special attack each turn")
	fmt.Println("4. Wild Pokemon will use basic attack (75% chance) or special attack (25% chance)")
	fmt.Println("5. Turns continue until one Pokemon faints")
	
	fmt.Println("\nBATTLE MECHANICS:")
	fmt.Println("- Basic attacks always hit")
	fmt.Println("- Special attacks have a 50% chance to hit")
	fmt.Println("- Special attacks can only be used ONCE per battle")
	
	fmt.Println("\nDAMAGE FORMULAS:")
	fmt.Println("- Basic attack: (skill damage + attacker's attack - defender's defense) × 0.1")
	fmt.Println("- Special attack: ((skill damage + attacker's special-attack) × 1.5 - defender's defense) × 0.1")
	
	fmt.Println("\nWHEN BATTLES OCCUR:")
	fmt.Println("- Battles are triggered when a Pokemon escapes during catch attempts")
	fmt.Println("- You'll have the option to battle or go back if a Pokemon escapes")
	fmt.Println("- If you win a battle, you'll catch the Pokemon")
	fmt.Println("- If you lose a battle, the Pokemon escapes")
	
	fmt.Println("\nTIPS:")
	fmt.Println("- Save your special attack for the right moment")
	fmt.Println("- Pokemon with higher attack stats are better at using basic attacks")
	fmt.Println("- Pokemon with higher special-attack stats excel with special attacks")
	fmt.Println("============================")
}