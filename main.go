package main

import (
	"time"

	"github.com/sphax7/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient 	pokeapi.Client
	nextLocation 	*string
	prevLocation 	*string
	caughtPokemon	map[string]pokeapi.Pokemon 		
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)

}