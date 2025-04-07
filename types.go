package main

// Export public name to make it accessible from other packages
var _ = Pokemon{}

// Pokemon represents a Pokemon with its attributes and stats
type Pokemon struct {
	Name           string  `json:"name"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
	BaseExperience int     `json:"base_experience"`
}

// Stat represents a specific Pokemon attribute's metadata
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Stats represents a specific Pokemon attribute with its value
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

// Type represents a Pokemon type's metadata
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Types represents a Pokemon's type with its slot
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}