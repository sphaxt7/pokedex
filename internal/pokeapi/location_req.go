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
	
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("got a cache")
		locationResp := RespLocations{}
		if err := json.Unmarshal(data, &locationResp); err != nil {
			return RespLocations{}, err
		}
	
		return locationResp, nil		
	}
	fmt.Println("no cache")


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


	data, err = io.ReadAll(res.Body)
	if err != nil {
		return RespLocations{}, err
	}

	locationResp := RespLocations{}
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(fullURL, data)

	return locationResp, nil
} 




func (c *Client) GetLocationInfo(LocationName string) (LocationInfo, error) {
	fullURL := baseURL + "/location-area/" + LocationName
	
	
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("got a cache")
		MyLocationInfo := LocationInfo{}
		if err := json.Unmarshal(data, &MyLocationInfo); err != nil {
			return LocationInfo{}, err
		}
	
		return MyLocationInfo, nil		
	}
	fmt.Println("no cache")


	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationInfo{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return LocationInfo{}, fmt.Errorf("error: %v", res.StatusCode)
	}


	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationInfo{}, err
	}

	MyLocationInfo := LocationInfo{}
	if err := json.Unmarshal(data, &MyLocationInfo); err != nil {
		return LocationInfo{}, err
	}

	c.cache.Add(fullURL, data)

	return MyLocationInfo, nil
} 