package main

import (
	"time"

	"github.com/sphax7/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient 	pokeapi.Client
	nextLocation 	*string
	prevLocation 	*string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&cfg)

}