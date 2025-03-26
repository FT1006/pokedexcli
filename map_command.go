package main

import (
	"fmt"
	"log"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *Config) error {
	areas, err := c.pokeapiClient.GetLocationAreaList(c.next)
	if err != nil {
		log.Fatal(err)
	}
	c.next = areas.Next
	c.prev = areas.Previous
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(c *Config) error {
	if c.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	areas, err := c.pokeapiClient.GetLocationAreaList(c.prev)
	if err != nil {
		log.Fatal(err)
	}
	c.next = areas.Next
	c.prev = areas.Previous
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}
