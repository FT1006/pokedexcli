package pokeapi

import (
	"net/http"
	"time"

	"github.com/FT1006/pokedexcli/internal/pokecache"
)

type Client struct {
	pokeapiClient *http.Client
	pokecache     *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		pokeapiClient: &http.Client{},
		pokecache:     pokecache.NewCache(time.Second * 5),
	}
}
