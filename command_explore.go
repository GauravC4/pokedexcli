package main

import (
	"fmt"

	"github.com/GauravC4/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("area name not found")
	}
	area := args[0]
	address := cfg.BaseURL + "/location-area/" + area
	exploreResp := pokeapi.ExploreResp{}
	err := pokeapi.Http_get(address, &exploreResp, cfg.CachePtr)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")

	for _, encounters := range exploreResp.PokemonEncounters {
		fmt.Printf("- %s\n", encounters.Pokemon.Name)
	}

	return nil
}
