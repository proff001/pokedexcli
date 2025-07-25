package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	desc string
	cb   func() error
}

var commands = map[string]cliCommand{}

func setCommands() {
	commands = map[string]cliCommand{
		"exit": {
			name: "exit",
			desc: "Exit the Pokedex",
			cb: commandExit,
		},
		"help": {
			name: "help",
			desc: "Displays a help message",
			cb: commandHelp,
		},
		"map": {
			name: "map",
			desc: "Displays a list of all the Pokemon locations",
			cb: commandMap,
		},
	}
}

func startRepl() {
	if len(commands) == 0 {
		setCommands()
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		rawInput := scanner.Text()
		input := cleanInput(rawInput)
		command := input[0]

		err := commands[command].cb()

		if err != nil {
			fmt.Print(err)
			continue
		}

		fmt.Print("\n")
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.desc)
	}

	return nil
}

func commandMap() error {
	locations, err := getLocations()
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Printf("%s\n", location)
	}

	return nil
}
