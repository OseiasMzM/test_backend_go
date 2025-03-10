package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Weather struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
		State   string `json:"state,omitempty"`
	} `json:"sys"`
}

const apiKey = "2ca18f18442db935049c4da5f5830d71"

func FetchWeatherByCity(city string) (Weather, error) {
	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", encodedCity, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("Error weather: status code %d", resp.StatusCode)
	}

	var weather Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return Weather{}, err
	}

	return weather, nil
}
