package pokeapi

import (
	"testing"
)

func TestBaseUrl(t *testing.T) {
	expected := "https://pokeapi.co/api/v2"
	if baseUrl != expected {
		t.Errorf("Expected baseUrl to be '%s', got '%s'", expected, baseUrl)
	}
}
