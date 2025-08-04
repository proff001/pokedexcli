package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/proff001/pokedexcli/internal/pokeapi"
)

type replConfig struct {
	pokeapiClient   pokeapi.Client
	locationNextUrl *string
	locationPrevUrl *string
	caughtPokemons  map[string]pokeapi.PokemonData
}

type replCommand struct {
	name string
	desc string
	cb   func(*replConfig, ...string) error
}

func startRepl(cfg *replConfig) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")

		scanner.Scan()
		rawInput := scanner.Text()
		input := cleanInput(rawInput)

		requestedcommand := input[0]
		args := input[1:]

		command, exists := getCommands()[requestedcommand]

		if !exists {
			fmt.Printf("Unknown command '%s'\n", requestedcommand)
			continue
		}

		err := command.cb(cfg, args...)

		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

func getCommands() map[string]replCommand {
	return map[string]replCommand{
		"exit": {
			name: "exit",
			desc: "Exit the Pokedex",
			cb:   commandExit,
		},
		"help": {
			name: "help",
			desc: "Displays a help message",
			cb:   commandHelp,
		},
		"map": {
			name: "map",
			desc: "Displays a page of 20 Pokemon locations, use 'mapb' to go back",
			cb:   commandMapf,
		},
		"mapb": {
			name: "mapb",
			desc: "Displays the previous page of 20 Pokemon locations",
			cb:   commandMapb,
		},
		"explore": {
			name: "explore <location name>",
			desc: "Displays all Pokemons you can encounter in a location",
			cb:   commandExplore,
		},
		"catch": {
			name: "catch <pokemon name>",
			desc: "Catch a Pokemon",
			cb:   commandCatch,
		},
		"inspect": {
			name: "inspect <pokemon name>",
			desc: "Inspect a Pokemon you have caught",
			cb:   inspectCommand,
		},
		"pokedex": {
			name: "pokedex",
			desc: "Displays a list of the Pokemon you have caught",
			cb:   pokedexCommand,
		},
	}
}
