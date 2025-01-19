package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/GauravC4/pokedexcli/internal/pokeapi"
	"github.com/GauravC4/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args []string) error
}

type config struct {
	BaseURL      string
	NextLocation string
	PrevLocation string
	Cache        pokecache.Cache
	Pokedex      pokeapi.Pokedex
}

var cfg = config{
	BaseURL:      "https://pokeapi.co/api/v2",
	NextLocation: "https://pokeapi.co/api/v2/location-area",
	PrevLocation: "",
	Cache:        pokecache.NewInMemoryCache(time.Minute * 5),
	Pokedex:      pokeapi.Pokedex{},
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display names of next 20 locations.",
			callback:    commandMapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "Display names of previous 20 locations.",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Accepts 1 argument <area-name> and displays pokemon encounters in that area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Accepts 1 argument <pokemon-name> and tries to catch it.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Accepts 1 argument <pokemon-name> and shows its stats if caught.",
			callback:    commandInspect,
		},
	}
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			return
		}

		cleanUserInputArr := CleanInput(scanner.Text())
		if len(cleanUserInputArr) == 0 {
			continue
		}

		command := cleanUserInputArr[0]
		args := cleanUserInputArr[1:]
		if cmd, ok := GetCommands()[command]; ok {
			err := cmd.callback(&cfg, args)
			if err != nil {
				fmt.Printf("Error : %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
