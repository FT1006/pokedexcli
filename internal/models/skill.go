package models

// Skill represents a move that a Pokemon can use
type Skill struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Damage int    `json:"damage"`
	Type   string `json:"type"`    // fire, water, etc.
	Class  string `json:"class"`   // physical, special, status
}