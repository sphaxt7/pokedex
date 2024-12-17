package main

import (
	// "math/rand"
	"errors"
	"fmt"
)

func commandInspect (cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon specified for inspection")
	}

	pokemonToInspect := args[0]
	pokemon, caught := cfg.caughtPokemon[pokemonToInspect]
	if !caught {
		return fmt.Errorf("you have not caught %s", pokemonToInspect)
	}


	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Species: %s\n", pokemon.Species.Name)
	// fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, typ := range pokemon.Types {
		fmt.Printf("\t- %s", typ.Type.Name)
	}

	return nil
}