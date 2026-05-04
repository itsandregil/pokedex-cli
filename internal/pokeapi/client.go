package pokeapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) ListLocations(pageURL *string) (LocationArea, error) {
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
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

	var locationArea LocationArea
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationArea); err != nil {
		return LocationArea{}, nil
	}

	return locationArea, nil
}
