package pokeapi

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func readAndUnmarshalJSON(filePath string) (Location, error) {
	var result Location // Variable to hold the unmarshaled data

	// 2. Read the file content
	jsonData, err := os.ReadFile(filePath) // Use ioutil.ReadFile for older Go
	if err != nil {
		return result, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	// 3. Unmarshal the JSON data into the struct
	err = json.Unmarshal(jsonData, &result) // Pass a pointer to the struct
	if err != nil {
		return result, fmt.Errorf("error unmarshaling JSON from %s: %w", filePath, err)
	}

	return result, nil
}

func TestStructJSON(T *testing.T) {
	filePath := "test_assets/types_location.json"
	data, err := readAndUnmarshalJSON(filePath)
	if err != nil {
		T.Fatalf("Failed to read/unmarshal %v", err)
	}

	T.Log(data)

}
