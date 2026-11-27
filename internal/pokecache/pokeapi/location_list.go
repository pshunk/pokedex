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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationArea{}, err
	}

	locations := locationArea{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return locationArea{}, err
	}

	return locations, nil
}
