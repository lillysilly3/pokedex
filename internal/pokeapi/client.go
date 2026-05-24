package pokeapi

import (
	"pokedex/internal/pokecache"
	"time"
	"net/http"
)

type Client struct {
	httpClient	http.Client
	cache		*pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	client := http.Client{
		Timeout: timeout,
	}
	cache := pokecache.NewCache(cacheInterval)
	return Client{
		httpClient: client,
		cache: cache,
	}
}
