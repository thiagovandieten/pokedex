package pokeapi

import "fmt"

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

func (p Pokemon) PrintInfo() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range p.Stats {
		fmt.Printf("\t -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Stats:\n")
	for _, t := range p.Types {
		fmt.Printf("\t- %s\n", t.Type.Name)
	}

}
