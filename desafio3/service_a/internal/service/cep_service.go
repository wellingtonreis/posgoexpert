package service

import (
	"encoding/json"
	"errors"
	"fmt"
	domain "service_a/internal/domain"

	fiber "github.com/gofiber/fiber/v2"
)

type ServiceCepImpl struct{}

func (s ServiceCepImpl) ServiceCep(number string) (domain.TemperatureLocation, error) {

	agent := fiber.Get("http://serviceb:9000/api/v1/cep/" + number + "/get")
	statusCode, body, errs := agent.Bytes()
	if errs != nil {
		return domain.TemperatureLocation{}, errors.New("failed to request")
	}

	if statusCode != fiber.StatusOK {
		return domain.TemperatureLocation{}, errors.New("invalid response")
	}

	var temperatureLocation domain.TemperatureLocation
	if err := json.Unmarshal(body, &temperatureLocation); err != nil {
		return domain.TemperatureLocation{}, fmt.Errorf("failed to parse temperature location: %w", err)
	}

	return temperatureLocation, nil
}
