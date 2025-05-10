package pokeapi

type Pokemon struct {
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Name           string  `json:"name"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
	Weight         int     `json:"weight"`
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

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
