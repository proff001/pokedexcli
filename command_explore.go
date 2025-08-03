package main

import "fmt"

func commandExplore(cfg *replConfig, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to specify a location")
	}

	locationData, err := cfg.pokeapiClient.GetLocationData(&args[0])

	if err != nil {
		return err
	}

	for _, enconters := range locationData.PokemonEncounters {
		fmt.Printf("%s\n", enconters.Pokemon.Name)
	}

	return nil
}
