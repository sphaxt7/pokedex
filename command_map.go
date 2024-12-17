package main

import (
	"errors"
	"fmt"
	// "github.com/sphax7/pokedex/internal/pokeapi"
)

func commandMapf(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.GetLocations(cfg.nextLocation)
	if err != nil {
		return err
	}

	fmt.Println("Locations Areas: ")
	for _, area := range resp.Results {
		fmt.Printf("-> %v \n", area.Name)
	}
	cfg.nextLocation = resp.Next
	cfg.prevLocation = resp.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocation == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.pokeapiClient.GetLocations(cfg.prevLocation)
	if err != nil {
		return err
	}

	fmt.Println("Locations Areas: ")
	for _, area := range resp.Results {
		fmt.Printf("-> %v \n", area.Name)
	}
	cfg.nextLocation = resp.Next
	cfg.prevLocation = resp.Previous

	return nil
}