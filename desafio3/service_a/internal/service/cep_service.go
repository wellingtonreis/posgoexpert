package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	domain "service_a/internal/domain"

	fiber "github.com/gofiber/fiber/v2"
)

type ServiceCepImpl struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (s ServiceCepImpl) ServiceCep(number string) (domain.TemperatureLocation, error) {

	if s.HTTPClient == nil {
		s.HTTPClient = &http.Client{}
	}

	url := fmt.Sprintf("%s/api/v1/cep/%s/get", s.BaseURL, number)

	resp, errClient := s.HTTPClient.Get(url)
	if errClient != nil {
		return domain.TemperatureLocation{}, fmt.Errorf("failed to get temperature location: %w", errClient)
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusOK {
		return domain.TemperatureLocation{}, errors.New("invalid response")
	}

	body, errResponse := io.ReadAll(resp.Body)
	if errResponse != nil {
		return domain.TemperatureLocation{}, fmt.Errorf("failed to read response body: %w", errResponse)
	}

	var temperatureLocation domain.TemperatureLocation
	if err := json.Unmarshal(body, &temperatureLocation); err != nil {
		return domain.TemperatureLocation{}, fmt.Errorf("failed to parse temperature location: %w", err)
	}

	return temperatureLocation, nil
}
