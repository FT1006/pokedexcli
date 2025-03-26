package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func MapCommand(c *Config) error {
	var err error
	var res *http.Response
	if c.Next != "" {
		res, err = http.Get(c.Next)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		res, err = http.Get("https://pokeapi.co/api/v2/location-area/")
		if err != nil {
			log.Fatal(err)
		}
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var areas LocationArea
	err = json.Unmarshal(body, &areas)
	if err != nil {
		log.Fatal(err)
	}
	c.Next = areas.Next
	c.Prev = areas.Previous
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func MapBackCommand(c *Config) error {
	var err error
	var res *http.Response
	if c.Prev != "" {
		res, err = http.Get(c.Prev)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var areas LocationArea
	err = json.Unmarshal(body, &areas)
	if err != nil {
		log.Fatal(err)
	}
	c.Next = areas.Next
	c.Prev = areas.Previous
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}
