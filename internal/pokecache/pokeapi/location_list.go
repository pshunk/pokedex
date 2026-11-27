package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (locationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locations := locationArea{}

	if bytes, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(bytes, &locations)
		return locations, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locations, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locations, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locations, err
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return locations, err
	}

	c.cache.Add(url, data)

	return locations, err
}
