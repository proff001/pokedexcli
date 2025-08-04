package main

import "fmt"

func pokedexCommand(cfg *replConfig, args ...string) error {
	if len(cfg.caughtPokemons) == 0 {
		return fmt.Errorf("you haven't caught any Pokemon yet")
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
