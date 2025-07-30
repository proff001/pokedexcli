package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(pageUrl *string) (LocationAreas, error) {
	url := BaseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreas{}, err
	}

	defer res.Body.Close()

	var data LocationAreas
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&data); err != nil {
		return LocationAreas{}, err
	}

	return data, nil
}
