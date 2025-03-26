package pokeapi

import "net/http"

type Client struct {
	pokeapiClient *http.Client
}

func NewClient() *Client {
	return &Client{
		pokeapiClient: &http.Client{},
	}
}
