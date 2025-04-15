package pokeapi

// MoveReference represents a basic reference to a move
type MoveReference struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// MoveDetails represents detailed information about a move
type MoveDetails struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Accuracy    int    `json:"accuracy"`
	Power       int    `json:"power"`
	PP          int    `json:"pp"`
	Priority    int    `json:"priority"`
	DamageClass struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"damage_class"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
	EffectChance int `json:"effect_chance"`
	// Additional fields that might be useful
	EffectEntries []struct {
		Effect      string `json:"effect"`
		ShortEffect string `json:"short_effect"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`
}

// PokemonMoves represents the moves a Pokemon can learn
type PokemonMoves struct {
	Move struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move"`
	VersionGroupDetails []struct {
		LevelLearnedAt  int `json:"level_learned_at"`
		VersionGroup    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
		MoveLearnMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move_learn_method"`
	} `json:"version_group_details"`
}