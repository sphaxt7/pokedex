package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func (c *Client) GetLocations(pageURL *string) (RespLocations, error) {
	fullURL := baseURL + "/location-area"
	
	if pageURL != nil {
		fullURL = *pageURL
	}
	
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return RespLocations{}, fmt.Errorf("error: %v", res.StatusCode)
	}


	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocations{}, err
	}

	locationResp := RespLocations{}
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return RespLocations{}, err
	}

	return locationResp, nil
} 