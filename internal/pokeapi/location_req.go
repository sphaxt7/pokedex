package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)



func GetLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	
	res, err := http.Get(url)
	if err != nil {
		return RespLocations{}, err
	}

	defer res.Body.Close()

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