package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, names ...string) error {
	if len(names) < 1 {
		return errors.New("no location name provided")
	}

	locationAreaName := names[0]
	locationArea, err := cfg.pokeapiClient.GetLocationInfo(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("---> Pokemon in: %s <---\n", locationArea.Name) 
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("=> %v \n", pokemon.Pokemon.Name)
	}

	return nil

}