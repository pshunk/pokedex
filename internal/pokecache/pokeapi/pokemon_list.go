package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName *string) (location, error) {
	location := location{}

	if locationName == nil {
		return location, nil
	}

	url := baseURL + "/location-area/" + *locationName

	if bytes, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(bytes, &location)
		return location, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return location, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return location, err
	}
	if resp.StatusCode > 299 {
		return location, fmt.Errorf("Could not find results for the given location")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return location, err
	}

	err = json.Unmarshal(data, &location)
	if err != nil {
		return location, err
	}

	c.cache.Add(url, data)

	return location, err
}
