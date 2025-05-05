package pokeapi

import (
	"testing"
)

func TestGetPokemonFromEterna(t *testing.T) {

	location, err := locationReadAndUnmarshalJSON("test_assets/eterna_forest.json")
	if err != nil {
		t.Fatalf("unable to read and/or unmarashal the provided file: %v ", err)
	}

	if len(location.PokemonEncounters) != 21 {
		t.Errorf("Didn't get the expected 21 pokemon from eterna")
	}
	// client := NewClient(5 * time.Second)

	// location, err := client.ListPokemon("eterna-forest")
	// if err != nil {
	// 	t.Errorf("Getting the pokemon from eterna forest went wrong")
	// }

}
