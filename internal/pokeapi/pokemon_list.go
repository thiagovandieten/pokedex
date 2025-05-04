package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(area string) (Location, error) {
	url := baseUrl + "/location-area/"
	if area != "" {
		url += area
	} else {
		return Location{}, errors.New("no area given")
	}

	//TODO: Check if in cache

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, nil
	}

	//TODO: Add to cache

	l := Location{}
	err = json.Unmarshal(dat, &l)
	if err != nil {
		return Location{}, err
	}

	return l, nil
}
