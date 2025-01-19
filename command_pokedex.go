package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.Pokedex) < 1 {
		return fmt.Errorf("no pokemon caught")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
