package main

import (
	"math/rand"
	"errors"
	"fmt"
	// "os"
)

func commandCatch (cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no target pokemon provided")
	}
	
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	fmt.Printf("%s was caught!!\n", pokemonName)

	return nil
}