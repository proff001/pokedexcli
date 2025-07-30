package pokeapi

import (
	"encoding/json"
	"io"
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

	cachedData, exists := c.cache.Get(url)

	if exists {
		data := LocationAreas{}
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return LocationAreas{}, err
		}

		return data, nil
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
	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationAreas{}, err
	}

	data := LocationAreas{}
	if err := json.Unmarshal(resBody, &data); err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(url, resBody)

	return data, nil
}
