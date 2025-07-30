package pokeapi

import (
	"net/http"
	"time"

	"github.com/proff001/pokedexcli/internal/pokecache"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{},
		cache:      pokecache.NewCache(cacheInterval),
	}
}
