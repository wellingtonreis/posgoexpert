package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	config "service_b/internal/config"
)

type ServiceWeatherApiImpl struct{}

type Weather struct {
	Location struct {
		Region string `json:"region"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	} `json:"current"`
}

func (s ServiceWeatherApiImpl) WeatherApi(q string) (Weather, error) {
	host := config.GetEnv("HOST_WEATHER_API", "http://api.weatherapi.com/v1/current.json")
	if host == "" {
		return Weather{}, errors.New("invalid host for WeatherAPI")
	}

	key := config.GetEnv("KEY_WEATHER_API", "32598b45dc044852846173447251801")
	if key == "" {
		return Weather{}, errors.New("invalid key for WeatherAPI")
	}

	urlApi := fmt.Sprintf("%s?key=%s&q=%s", host, key, url.QueryEscape(q))

	resp, err := http.Get(urlApi)
	if err != nil {
		return Weather{}, errors.New("failed to request WeatherAPI")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Weather{}, errors.New("invalid response from WeatherAPI")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Weather{}, fmt.Errorf("failed to read response from WeatherAPI: %w", err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return Weather{}, errors.New("failed to parse response from WeatherAPI")
	}

	return weather, nil
}
