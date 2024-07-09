package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/BatmiBoom/pokedex/cmd/cache"
)

const BaseURI = "https://pokeapi.co/api/v2/"
const LocationAreasURI = BaseURI + "location-area/"

type Client struct {
	cache cache.Cache
	http  http.Client
}

func NewClient(timeout, cache_interval time.Duration) Client {
	return Client{
		cache: cache.NewCache(timeout),
		http: http.Client{
			Timeout: timeout,
		},
	}
}

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var c = NewClient(5*time.Second, time.Minute*5)

func GetLocations(URI string) Locations {
	var locations Locations

	if val, ok := c.cache.Get(URI); ok {
		err := json.Unmarshal(val, &locations)
		if err != nil {
			fmt.Printf("error: there was an error unmarshilling the cache %v \n", err)
		}

		return locations
	}

	resp, err := c.http.Get(URI)
	if err != nil {
		fmt.Printf("error: there was an error getting the areas %v \n", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: there was an error with the response %v \n", err)
	}

	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Printf("error: there was an error converting json to go structs %v \n", err)
	}

	c.cache.Add(URI, body)
	return locations
}
