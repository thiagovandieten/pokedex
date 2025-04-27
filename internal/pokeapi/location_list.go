package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// func GetLocationAreas(query_param string) (LocationAreas, error) {
// 	var location_areas LocationAreas
// 	// fmt.Printf("LOG: apiURLq + query_param = %s%s\ne", apiURL, query_param)
// 	res, err := http.Get(apiURL + query_param)
// 	if err != nil {
// 		return location_areas, fmt.Errorf("REQUEST FAILED: %v", err)
// 	}

// 	decoder := json.NewDecoder(res.Body)
// 	err = decoder.Decode(&location_areas)
// 	if err != nil {
// 		return location_areas, fmt.Errorf("DECODING THE BODY FAILED:: %v", err)
// 	}
// 	return location_areas, nil
// }

func (c *Client) ListLocations(pageURL *string) (LocationAreas, error) {
	url := baseUrl + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if value, ok := c.cache.GetCache(url); ok {
		//TODO: It's the same code as 60-66. Should extract to own function
		la := LocationAreas{}
		err := json.Unmarshal(value, &la)
		if err != nil {
			return LocationAreas{}, err
		}

		return la, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	c.cache.AddCache(url, dat)

	la := LocationAreas{}
	err = json.Unmarshal(dat, &la)
	if err != nil {
		return LocationAreas{}, err
	}

	return la, nil
}
