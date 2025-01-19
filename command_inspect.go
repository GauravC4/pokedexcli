package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("pokemon name not found")
	}
	pokemon, ok := cfg.Pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	if len(pokemon.Stats) > 0 {
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
	}

	fmt.Println("Types:")
	for _, typeVal := range pokemon.Types {
		fmt.Printf(" - %s\n", typeVal.Type.Name)
	}
	return nil
}
