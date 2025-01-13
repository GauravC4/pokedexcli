package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, val := range GetCommands() {
		fmt.Printf("%v: %v\n", val.name, val.description)
	}
	return nil
}
