package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"
)


func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// Check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// Found the data in the cache!
		fmt.Println("BOOM! data found in cache!")
		locationAreasResp := LocationAreasResponse{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil{
			return LocationAreasResponse{}, err
		}
		return locationAreasResp, nil
	}
	fmt.Println("No Cache for you!")


	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil{
		return LocationAreasResponse{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil{
		return LocationAreasResponse{}, err
	}

	locationAreasResp := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil{
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}


func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// Check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// Found the data in the cache!
		fmt.Println("BOOM! data found in cache!")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil{
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	fmt.Println("No Cache for you!")


	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil{
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil{
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil{
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationArea, nil
}

