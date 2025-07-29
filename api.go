package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocations(url string) (LocationAreas, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	res, err := http.Get(url)

	if err != nil {
		return LocationAreas{}, fmt.Errorf("error fetching locations: %s", err)
	}

	defer res.Body.Close()

	var data any
	locations := LocationAreas{}
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&data); err != nil {
		return LocationAreas{}, fmt.Errorf("error decoding locations: %s", err)
	}

	return locations, nil
}