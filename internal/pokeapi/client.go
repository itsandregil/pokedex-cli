package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/itsandregil/pokedex-cli/internal/cache"
)

type Client struct {
	httpClient http.Client
	cache      cache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache.NewCache(cacheInterval),
	}
}

func (c *Client) ListLocations(pageURL *string) (LocationArea, error) {
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cached, ok := c.cache.Get(url); ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(cached, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)
	return locationArea, nil
}
