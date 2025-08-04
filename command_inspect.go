package main

import "fmt"

func inspectCommand(cfg *replConfig, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to specify a Pokemon")
	}

	pokemonData, ok := cfg.caughtPokemons[args[0]]

	if !ok {
		return fmt.Errorf("you haven't caught %s yet", args[0])
	}

	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemonData.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, typeData := range pokemonData.Types {
		fmt.Printf(" - %s\n", typeData.Type.Name)
	}

	return nil
}
