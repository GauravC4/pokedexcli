package main

import (
	"fmt"

	"github.com/GauravC4/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config, back bool) error {
	address := cfg.NextLocation
	if back {
		if len(cfg.PrevLocation) == 0 {
			fmt.Println("you're on the first page")
			return nil
		}
		address = cfg.PrevLocation
	}
	locationResp := pokeapi.LocationResp{}
	err := pokeapi.Http_get(address, &locationResp, cfg.Cache)
	if err != nil {
		return err
	}

	cfg.NextLocation = locationResp.Next
	cfg.PrevLocation = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapNext(cfg *config, args []string) error {
	return commandMap(cfg, false)
}

func commandMapBack(cfg *config, args []string) error {
	return commandMap(cfg, true)
}
