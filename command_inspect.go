package main

import "fmt"

func commandInspect(pokemon string, c *Config) error {
	if pokemon == "" {
		fmt.Println("no pokemon provided")
		return nil
	}
	if pkm, ok := c.caughtPokemon[pokemon]; !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", pkm.Name)
		fmt.Printf("Height: %d\n", pkm.Height)
		fmt.Printf("Weight: %d\n", pkm.Weight)
		fmt.Println("Stats:")
		for _, stat := range pkm.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typ := range pkm.Types {
			fmt.Printf("  -%s\n", typ.Type.Name)
		}
	}
	return nil
}
