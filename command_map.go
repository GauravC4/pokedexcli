package main

import (
	"fmt"

	"github.com/GauravC4/pokedexcli/internal/pokeapi"
	"github.com/GauravC4/pokedexcli/internal/utils"
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
	err := utils.Http_get(address, &locationResp)
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

func commandMapNext(cfg *config) error {
	return commandMap(cfg, false)
}

func commandMapBack(cfg *config) error {
	return commandMap(cfg, true)
}
