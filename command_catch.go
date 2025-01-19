package main

import (
	"fmt"
	"math/rand"

	"github.com/GauravC4/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("pokemon name not found")
	}
	pokemon := args[0]
	address := cfg.BaseURL + "/pokemon/" + pokemon
	pokemonResp := pokeapi.Pokemon{}
	err := pokeapi.Http_get(address, &pokemonResp, cfg.Cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonResp.Name)
	if rand.Intn(400) > pokemonResp.BaseExperience {
		fmt.Printf("%v was caught!\n", pokemonResp.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.Pokedex[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Printf("%v escaped!\n", pokemonResp.Name)
	}
	return nil
}
