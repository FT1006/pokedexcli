package pokeapi

type PokemonRes struct {
	Name           string        `json:"name"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []Stats       `json:"stats"`
	Types          []Types       `json:"types"`
	BaseExperience int           `json:"base_experience"`
	Moves          []PokemonMove `json:"moves"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

// PokemonMove represents a move a Pokemon can learn
type PokemonMove struct {
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
