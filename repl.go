package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			err := cmd.callback()
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
