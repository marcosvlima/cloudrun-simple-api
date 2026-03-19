package weatherapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type WeatherAPIClient struct {
	APIKey string
}

func NewWeatherAPIClient(apiKey string) *WeatherAPIClient {
	return &WeatherAPIClient{APIKey: apiKey}
}

func (c *WeatherAPIClient) GetTemperature(city string) (float64, error) {
	reqURL := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", c.APIKey, url.QueryEscape(city))
	resp, err := http.Get(reqURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("weatherapi error: status %d", resp.StatusCode)
	}

	var data WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	return data.Current.TempC, nil
}
