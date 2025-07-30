package main

import (
	"time"

	"github.com/proff001/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 3*time.Minute)

	config := &replConfig{
		pokeapiClient: client,
	}

	startRepl(config)
}
