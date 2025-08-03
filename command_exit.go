package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *replConfig, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
