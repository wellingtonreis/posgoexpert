package usecase

import (
	domain "service_b/internal/domain"
	service "service_b/internal/service"
)

type CepService interface {
	ViaCep(number string) (service.Location, error)
}

type WeatherApi interface {
	WeatherApi(q string) (service.Weather, error)
}

type TemperatureRecoveryLocationUseCase struct {
	cepService  CepService
	wheatherApi WeatherApi
}

func NewTemperatureRecoveryLocationUseCase(cepService CepService, weatherApi WeatherApi) *TemperatureRecoveryLocationUseCase {
	return &TemperatureRecoveryLocationUseCase{
		cepService:  cepService,
		wheatherApi: weatherApi,
	}
}

func (u *TemperatureRecoveryLocationUseCase) GetTemperatureRecoveryLocation(number string) (domain.TemperatureLocation, error) {
	location, err := u.cepService.ViaCep(number)
	if err != nil {
		return domain.TemperatureLocation{}, err
	}

	weather, err := u.wheatherApi.WeatherApi(location.Localidade)
	if err != nil {
		return domain.TemperatureLocation{}, err
	}

	celsius := float32(weather.Current.TempC)
	farenheit := celsius*1.8 + 32
	Kelvin := celsius + 273.15

	temperatureLocation := domain.TemperatureLocation{
		Number:     location.Cep,
		City:       location.Localidade,
		Celsius:    celsius,
		Fahrenheit: farenheit,
		Kelvin:     Kelvin,
	}

	return temperatureLocation, nil
}
