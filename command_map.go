package main

import "fmt"

func commandMapf(cfg *replConfig) error {
	locationData, err := cfg.pokeapiClient.GetLocations(cfg.locationNextUrl)

	if err != nil {
		return err
	}

	cfg.locationNextUrl = locationData.Next
	cfg.locationPrevUrl = locationData.Previous

	for _, location := range locationData.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func commandMapb(cfg *replConfig) error {
	if cfg.locationPrevUrl == nil {
		return fmt.Errorf("you're on the first page")
	}

	locationData, err := cfg.pokeapiClient.GetLocations(cfg.locationPrevUrl)

	if err != nil {
		return err
	}

	cfg.locationNextUrl = locationData.Next
	cfg.locationPrevUrl = locationData.Previous

	for _, location := range locationData.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
