package main

import (
	"time"

	"github.com/proff001/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)

	config := &replConfig{
		pokeapiClient: client,
	}

	startRepl(config)
}
