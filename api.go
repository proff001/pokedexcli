package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const url = "https://pokeapi.co/api/v2/"

func getLocations() ([]string, error) {
	endpoint := "location-area"
	res, err := http.Get(url + endpoint)

	if err != nil {
		return nil, fmt.Errorf("error fetching locations: %s", err)
	}

	defer res.Body.Close()

	var data any
	locations := []string{}
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding locations: %s", err)
	}

	fmt.Printf("%+v\n", data)
	
	return locations, nil
}