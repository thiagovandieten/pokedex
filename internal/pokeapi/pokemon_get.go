package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/"
	if name != "string" {
		url += name
	} else {
		return Pokemon{}, errors.New("no name of pokemon given")
	}

	if value, ok := c.cache.GetCache(url); ok {
		p := Pokemon{}
		err := json.Unmarshal(value, &p)
		if err != nil {
			return Pokemon{}, err
		}
		return p, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.AddCache(url, dat)

	p := Pokemon{}
	err = json.Unmarshal(dat, &p)
	if err != nil {
		return Pokemon{}, err
	}
	return p, nil
}
