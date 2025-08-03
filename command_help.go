package main

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *replConfig, args ...string) error {
	commands := getCommands()
	keys := make([]string, 0, len(commands))

	for key := range commands {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, commandKey := range keys {
		command := commands[commandKey]
		fmt.Printf("%s: %s\n", command.name, command.desc)
	}

	return nil
}
