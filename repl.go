package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			return
		}
		userInput := scanner.Text()
		cleanUserInputArr := CleanInput(userInput)
		if len(cleanUserInputArr) > 0 {
			fmt.Printf("Your command was: %v\n", cleanUserInputArr[0])
		}
	}
}

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
