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

	if value, ok := c.cache.GetCache(url); ok {
		l := Location{}
		err := json.Unmarshal(value, &l)
		if err != nil {
			return Location{}, err
		}
		return l, nil
	}

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

	c.cache.AddCache(url, dat)

	l := Location{}
	err = json.Unmarshal(dat, &l)
	if err != nil {
		return Location{}, err
	}

	return l, nil
}
