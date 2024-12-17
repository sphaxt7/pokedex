package pokeapi

import (
	"net/http"
	"time"

	pokecache "github.com/sphax7/pokedex/internal/pokeache"
)


const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}


func NewClient(cacheInterval time.Duration) Client {
	return Client {
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}


