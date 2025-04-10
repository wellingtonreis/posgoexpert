package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	domain "service_a/internal/domain"
)

type ServiceCepImpl struct {
	HTTPClient *http.Client
	BaseURL    string
}

func NewServiceCepImpl(httpClient *http.Client, baseURL string) *ServiceCepImpl {

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &ServiceCepImpl{
		HTTPClient: httpClient,
		BaseURL:    baseURL,
	}
}

func (s *ServiceCepImpl) ServiceCep(number string) (domain.TemperatureLocation, error) {

	if s.HTTPClient == nil {
		return domain.TemperatureLocation{}, errors.New("http client is not initialized")
	}

	url := fmt.Sprintf("%s/api/v1/cep/%s/get", s.BaseURL, number)

	resp, errClient := s.HTTPClient.Get(url)
	if errClient != nil {
		return domain.TemperatureLocation{}, fmt.Errorf("failed to get temperature location: %w", errClient)
	}
	defer resp.Body.Close()

	body, errResponse := io.ReadAll(resp.Body)
	if errResponse != nil {
		return domain.TemperatureLocation{}, fmt.Errorf("failed to read response body: %w", errResponse)
	}

	var jsonResponse struct {
		Error string `json:"error"`
	}

	if err := json.Unmarshal(body, &jsonResponse); err == nil && jsonResponse.Error != "" {
		return domain.TemperatureLocation{}, fmt.Errorf("%s", jsonResponse.Error)
	}

	var temperatureLocation domain.TemperatureLocation
	if err := json.Unmarshal(body, &temperatureLocation); err != nil {
		return domain.TemperatureLocation{}, fmt.Errorf("failed to parse temperature location: %w", err)
	}

	return temperatureLocation, nil
}
