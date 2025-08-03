package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *replConfig, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to specify a Pokemon")
	}

	pokemonData, err := cfg.pokeapiClient.GetPokemonData(&args[0])

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonData.Name)

	decay := 0.0052
	min := 0.01
	max := 0.99
	chance := min + (max-min)*math.Exp(-decay*float64(pokemonData.BaseExperience))

	if chance > max {
		chance = max
	}

	if chance < min {
		chance = min
	}

	if rand.Float64() < chance {
		fmt.Printf("%s was caught!\n", pokemonData.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
	}

	cfg.caughtPokemons[pokemonData.Name] = pokemonData

	return nil
}
