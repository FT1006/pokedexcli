package main

import (
	"fmt"
	"log"
)

func commandMap(additionalInput string, c *Config) error {
	if additionalInput != "" {
		fmt.Println("Additional input has been ignored")
	}
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

func commandMapBack(additionalInput string, c *Config) error {
	if additionalInput != "" {
		fmt.Println("Additional input has been ignored")
	}
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
