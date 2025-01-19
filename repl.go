package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/GauravC4/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	BaseURL      string
	NextLocation string
	PrevLocation string
	CachePtr     *(pokecache.Cache)
}

var cfg = config{
	BaseURL:      "https://pokeapi.co/api/v2",
	NextLocation: "https://pokeapi.co/api/v2/location",
	PrevLocation: "",
	CachePtr:     pokecache.NewCache(time.Second * 5),
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "display names of next 20 locations",
			callback:    commandMapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "display names of previous 20 locations",
			callback:    commandMapBack,
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
		if cmd, ok := GetCommands()[command]; ok {
			err := cmd.callback(&cfg)
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
